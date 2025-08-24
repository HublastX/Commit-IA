#!/usr/bin/env node
const { execSync, spawnSync } = require("child_process");
const os = require("os");
const path = require("path");
const fs = require("fs");

const platform = os.platform();
const arch = os.arch();
const version = "v2.0.0";
const repo = "HublastX/Commit-IA";

// Mapear arquiteturas
const archMap = {
  "x64": "amd64",
  "arm64": "arm64"
};

// Mapear plataformas
const platformMap = {
  "linux": "linux",
  "darwin": "darwin",
  "win32": "windows"
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

// Para teste, vamos focar só no Linux primeiro
if (platform !== "linux") {
  console.log("⚠️  Por enquanto, apenas Linux é suportado para testes.");
  console.log("   Em breve: macOS e Windows");
  process.exit(1);
}

const binName = platform === "win32" ? "commitia.exe" : "commitia";
const binPath = path.join(distDir, binName);

// URL do release - usando o nome atual da sua release
const fileName = `commitia`; // Seu binário atual é só "commitia"
const url = `https://github.com/${repo}/releases/download/${version}/${fileName}`;

console.log(`📦 Instalando CommitIA ${version} para ${platform}/${arch}`);

function downloadBinary() {
  try {
    console.log(`⬇️  Baixando de: ${url}`);
    
    // Usar curl com flags mais robustas
    const result = spawnSync("curl", [
      "-f",           // Falha silenciosamente em erros HTTP
      "-L",           // Segue redirects
      "-S",           // Mostra erros
      "--progress-bar", // Barra de progresso
      "-o", binPath,  // Output para arquivo
      url
    ], { 
      stdio: ["inherit", "inherit", "pipe"],
      encoding: "utf8"
    });

    if (result.status !== 0) {
      throw new Error(`Download falhou com código ${result.status}: ${result.stderr}`);
    }

    // Verificar se arquivo foi baixado
    if (!fs.existsSync(binPath) || fs.statSync(binPath).size === 0) {
      throw new Error("Arquivo baixado está vazio ou não existe");
    }

    // Tornar executável no Linux/macOS
    if (platform !== "win32") {
      fs.chmodSync(binPath, 0o755);
    }

    console.log(`✅ CommitIA instalado com sucesso!`);
    console.log(`   Binário: ${binPath}`);
    console.log(`   Teste com: npx commitia --help`);

  } catch (error) {
    console.error(`❌ Erro no download: ${error.message}`);
    
    // Verificar se já existe binário local
    if (fs.existsSync(binPath) && fs.statSync(binPath).size > 0) {
      console.log(`⚠️  Usando binário local existente em ${binPath}`);
      if (platform !== "win32") {
        fs.chmodSync(binPath, 0o755);
      }
      return;
    }

    console.error("\n💡 Soluções possíveis:");
    console.error("1. Verifique sua conexão com a internet");
    console.error("2. Confirme se a release existe no GitHub");
    console.error("3. Compile manualmente:");
    console.error(`   GOOS=${mappedPlatform} GOARCH=${mappedArch} go build -o dist/${binName}`);
    process.exit(1);
  }
}

// Verificar dependências
function checkDependencies() {
  try {
    spawnSync("curl", ["--version"], { stdio: "pipe" });
  } catch (error) {
    console.error("❌ curl não encontrado. Instale com:");
    console.error("   Ubuntu/Debian: sudo apt install curl");
    console.error("   CentOS/RHEL: sudo yum install curl");
    process.exit(1);
  }
}

// Main
checkDependencies();
downloadBinary();