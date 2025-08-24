#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");
const os = require("os");

const platform = os.platform();
const binName = platform === "win32" ? "commitai.exe" : "commitai";
const binPath = path.join(__dirname, "..", "dist", binName);

function runcommitai() {
    if (!fs.existsSync(binPath)) {
        console.error("❌ Binary not found!");
        console.error(`   Expected: ${binPath}`);
        console.error("\n💡 Run: npm run postinstall");
        process.exit(1);
    }

    try {
        fs.accessSync(binPath, fs.constants.F_OK | fs.constants.X_OK);
    } catch (error) {
        if (platform !== "win32") {
            try {
                fs.chmodSync(binPath, 0o755);
            } catch (chmodError) {
                console.error("❌ Permission error:", chmodError.message);
                process.exit(1);
            }
        }
    }

    const args = process.argv.slice(2);
    const child = spawn(binPath, args, {
        stdio: "inherit",
        shell: false,
    });

    child.on("error", (error) => {
        console.error("❌ Execution error:", error.message);
        process.exit(1);
    });

    child.on("exit", (code, signal) => {
        if (signal) {
            process.exit(1);
        }
        process.exit(code || 0);
    });
}

runcommitai();
