import fs from 'fs';
import path from 'path';

function findProjectRoot(startDir) {
  let currentDir = path.resolve(startDir);
  const root = path.parse(currentDir).root;

  while (currentDir !== root) {
    const markers = ['.git', '.agents', '.agent', 'package.json', 'pyproject.toml', 'Cargo.toml', 'go.mod', 'pom.xml'];
    for (const marker of markers) {
      if (fs.existsSync(path.join(currentDir, marker))) {
        return currentDir;
      }
    }
    currentDir = path.dirname(currentDir);
  }
  return null;
}

const cwd = process.cwd();
const projectRoot = findProjectRoot(cwd);

let memoryFile;
if (projectRoot) {
  const agentsDir = path.join(projectRoot, '.agents');
  if (fs.existsSync(agentsDir)) {
    memoryFile = path.join(agentsDir, 'memory.jsonl');
  } else {
    memoryFile = path.join(projectRoot, '.memory.jsonl');
  }
} else {
  const globalDir = path.join(process.env.USERPROFILE || process.env.HOME || '.', '.gemini', 'antigravity');
  if (!fs.existsSync(globalDir)) {
    fs.mkdirSync(globalDir, { recursive: true });
  }
  memoryFile = path.join(globalDir, 'global_memory.jsonl');
}

process.env.MEMORY_FILE_PATH = memoryFile;

// Dynamically import the real server-memory
const serverPath = 'file:///' + 'C:/Users/Admin/AppData/Roaming/npm/node_modules/@modelcontextprotocol/server-memory/dist/index.js'.replace(/\\/g, '/');
import(serverPath);
