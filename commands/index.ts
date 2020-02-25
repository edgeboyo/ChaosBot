import * as Discord from "discord.js";

// Add new commands to this file

import chaos from "./chaos";

const commands: {
	[command: string]: (message: Discord.Message, client: Discord.Client, args: string[]) => Promise<void>;
} = {
	chaos,
};

export default async function(
	command: string,
	message: Discord.Message,
	client: Discord.Client,
	args: string[]
): Promise<void> {
	await commands[command]?.(message, client, args);
}
