#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");
const os = require("os");

const platform = os.platform();
const binName = platform === "win32" ? "commitia.exe" : "commitia";
const binPath = path.join(__dirname, "..", "dist", binName);

function runCommitIA() {
  if (!fs.existsSync(binPath)) {
    console.error("❌ Binário não encontrado!");
    console.error(`   Esperado: ${binPath}`);
    console.error("\n💡 Execute: npm run postinstall");
    process.exit(1);
  }

  try {
    fs.accessSync(binPath, fs.constants.F_OK | fs.constants.X_OK);
  } catch (error) {
    if (platform !== "win32") {
      try {
        fs.chmodSync(binPath, 0o755);
      } catch (chmodError) {
        console.error("❌ Erro de permissões:", chmodError.message);
        process.exit(1);
      }
    }
  }

  const args = process.argv.slice(2);
  const child = spawn(binPath, args, {
    stdio: "inherit",
    shell: false
  });

  child.on("error", (error) => {
    console.error("❌ Erro ao executar:", error.message);
    process.exit(1);
  });

  child.on("exit", (code, signal) => {
    if (signal) {
      process.exit(1);
    }
    process.exit(code || 0);
  });
}

runCommitIA();