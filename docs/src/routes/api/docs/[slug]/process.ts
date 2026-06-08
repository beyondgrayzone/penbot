import { spawn } from "child_process";
import path from "path";
import { fileURLToPath } from "url";

export const runProcessMds = (
	targetPackageName: string,
	scriptName: string,
	inputPath: string,
	outputPath: string,
): Promise<string> => {
	return new Promise((resolve, reject) => {
		const currentFile = fileURLToPath(import.meta.url);
		const normalizedFile = currentFile.split(path.sep).join("/");

		let parts = normalizedFile.split("/docs/src/routes");

		if (parts.length < 2) {
			parts = normalizedFile.split("/docs/.svelte-kit");
		}

		if (parts.length < 2) {
			const errorMsg = `Error: The current file is not located inside a 'docs' directory.\nPath: ${currentFile}`;
			console.error(`normalizedFile => ${normalizedFile}`);
			console.error(`inputPath => ${inputPath}`);
			console.error(errorMsg);
			console.error(parts[0]);
			return reject(new Error(errorMsg));
		}

		const basePath = path.join(parts[0], "docs");
		const absoluteInputPath = path.join(basePath, inputPath);
		const absoluteOutputPath = path.join(basePath, outputPath);

		console.log(`> Triggering '${scriptName}' in '${targetPackageName}'...`);
		console.log(`> Context Root: ${basePath}`);
		console.log(`> Input: ${absoluteInputPath}`);
		console.log(`> Output: ${absoluteOutputPath}`);
		console.log(`==> ${currentFile}`);

		const pnpmCommand = "bun";

		const args = [
			"run",
			"--filter",
			targetPackageName,
			scriptName,
			"-in",
			absoluteInputPath,
			"-out",
			absoluteOutputPath,
		];

		const child = spawn(pnpmCommand, args, {
			stdio: "inherit",
			shell: true,
		});

		child.on("error", (err) => {
			console.error(`> Failed to start subprocess: ${err.message}`);
			reject(err);
		});

		child.on("close", (code: number) => {
			if (code === 0) {
				console.log(`\n> Success! Output generated at: ${absoluteOutputPath}`);
				resolve(absoluteOutputPath);
			} else {
				console.error(`\n> Process exited with code ${code}`);
				reject(new Error(`Child process failed with code ${code}`));
			}
		});
	});
};
