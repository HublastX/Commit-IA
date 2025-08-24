#!/usr/bin/env node
const { spawnSync } = require("child_process");
const os = require("os");
const path = require("path");
const fs = require("fs");

const platform = os.platform();
const arch = os.arch();
const version = "v2.0.0";
const repo = "HublastX/Commit-IA";

// Mapear arquiteturas e plataformas
const archMap = {
    x64: "amd64",
    arm64: "arm64",
};

const platformMap = {
    linux: "linux",
    darwin: "darwin",
    win32: "windows",
};

const distDir = path.join(__dirname, "..", "dist");
if (!fs.existsSync(distDir)) {
    fs.mkdirSync(distDir, { recursive: true });
}

const mappedPlatform = platformMap[platform];
const mappedArch = archMap[arch];

if (!mappedPlatform || !mappedArch) {
    console.error(`❌ Plataforma não suportada: ${platform}/${arch}`);
    process.exit(1);
}

const binName = platform === "win32" ? "commitai.exe" : "commitai";
const binPath = path.join(distDir, binName);

let fileName;
if (platform === "win32") {
    fileName = "commitai.exe";
} else {
    fileName = "commitai";
}

const url = `https://github.com/${repo}/releases/download/${version}/${fileName}`;

console.log(`📦 Instalando commitai ${version} para ${platform}/${arch}`);

function checkDependencies() {
    try {
        spawnSync("curl", ["--version"], { stdio: "pipe" });
    } catch (error) {
        console.error("❌ curl não encontrado. Instale curl primeiro.");
        process.exit(1);
    }
}

function downloadBinary() {
    try {
        console.log(`⬇️  Baixando: ${fileName}`);

        const result = spawnSync(
            "curl",
            ["-f", "-L", "-S", "--progress-bar", "-o", binPath, url],
            {
                stdio: ["inherit", "inherit", "pipe"],
                encoding: "utf8",
            }
        );

        if (result.status !== 0) {
            throw new Error(`Download falhou: ${result.stderr}`);
        }

        if (!fs.existsSync(binPath) || fs.statSync(binPath).size === 0) {
            throw new Error("Arquivo baixado está vazio");
        }

        if (platform !== "win32") {
            fs.chmodSync(binPath, 0o755);
        }

        console.log(`✅ commitai instalado com sucesso!`);
        console.log(`   Teste: npx commitai --help`);
    } catch (error) {
        console.error(`❌ Erro: ${error.message}`);
        console.error("\n💡 Verifique:");
        console.error("1. Conexão com internet");
        console.error("2. Se a release existe no GitHub");
        console.error(`3. URL: ${url}`);
        process.exit(1);
    }
}

checkDependencies();
downloadBinary();
