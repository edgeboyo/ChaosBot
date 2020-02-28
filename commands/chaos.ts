import * as Discord from "discord.js";

export default async function(message: Discord.Message, client: Discord.Client, args: string[]) {
	// makes the bot say something and delete the message. As an example, it's open to anyone to use.
	// To get the "message" itself we join the `args` back into a string with spaces:
	const sayMessage = args.join(" ");
	
	
	const backwds = "zʎxʍʌnʎsɹbdouɯʅʞɾᴉɥƃⅎǝpɔqɐ";
	const backup = "Z⅄XϺɅՈꓕSꓤꝹԀONꟽ⅂ꓘᒋIH⅁ᖵƎᗡϽꓭ∀";
	// Then we delete the command message (sneaky, right?). The catch just ignores the error with a cute smiley thing.
	message.delete().catch(O_o => {});

	var sendMessg = "";
	for(var i = 0; i < sayMessage.length; i++){
		var ch = sayMessage[i];

		if('a' <= ch && ch <= 'z'){
			var no = ch.charCodeAt(0) - 'a'.charCodeAt(0);
			ch = backwds[backwds.length - 1 - no];
		}
		else if ('A' <= ch && ch <= 'Z'){
			var no = ch.charCodeAt(0) - 'A'.charCodeAt(0);
			ch = backup[backup.length - 1 - no];
		}
		sendMessg = ch + sendMessg;
	}
	// And we get the bot to say the thing:
	message.channel.send(sendMessg);
}
