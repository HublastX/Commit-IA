#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");
const os = require("os");

const platform = os.platform();
const binName = platform === "win32" ? "commitia.exe" : "commitia";
const binPath = path.join(__dirname, "..", "dist", binName);

function runCommitIA() {
  // Verificar se o binário existe
  if (!fs.existsSync(binPath)) {
    console.error("❌ Binário CommitIA não encontrado!");
    console.error(`   Esperado em: ${binPath}`);
    console.error("\n💡 Tente executar:");
    console.error("   npm run postinstall");
    console.error("   ou");
    console.error("   node bin/install.js");
    process.exit(1);
  }

  // Verificar se é executável
  try {
    fs.accessSync(binPath, fs.constants.F_OK | fs.constants.X_OK);
  } catch (error) {
    console.error("❌ Binário encontrado mas não é executável!");
    if (platform !== "win32") {
      console.log("🔧 Tentando corrigir permissões...");
      try {
        fs.chmodSync(binPath, 0o755);
        console.log("✅ Permissões corrigidas!");
      } catch (chmodError) {
        console.error("❌ Não foi possível corrigir permissões:", chmodError.message);
        process.exit(1);
      }
    } else {
      process.exit(1);
    }
  }

  // Executar o binário diretamente - deixar o Go lidar com todos os casos
  const args = process.argv.slice(2);
  
  const child = spawn(binPath, args, {
    stdio: "inherit",
    shell: false
  });

  child.on("error", (error) => {
    if (error.code === "ENOENT") {
      console.error("❌ Não foi possível executar o binário CommitIA");
      console.error(`   Caminho: ${binPath}`);
      console.error("\n💡 Se o problema persistir, tente:");
      console.error("   npm run postinstall");
    } else {
      console.error("❌ Erro ao executar CommitIA:", error.message);
    }
    process.exit(1);
  });

  child.on("exit", (code, signal) => {
    if (signal) {
      console.error(`CommitIA foi terminado pelo sinal: ${signal}`);
      process.exit(1);
    }
    process.exit(code || 0);
  });
}

// Executar
runCommitIA();