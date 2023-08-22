from dotenv import dotenv_values
from aiogram import Bot, Dispatcher, executor, types

async def choose_language(message: types.Message):
		buttons = [
			[
				types.KeyboardButton(text='Русский', callback_data='ru'),
                types.KeyboardButton(text='English', callback_data='en'),
                types.KeyboardButton(text='Español', callback_data='es'),
			],
		]

		keyboard_markup = types.ReplyKeyboardMarkup(
			keyboard=buttons,
			resize_keyboard=True,
			input_field_placeholder="You can change it at any time in the settings"
		)
		
		await message.answer(text="Please, select a language", reply_markup=keyboard_markup)


async def main_menu_screen(message: types.Message):
	# TODO: normal database query for user balance
	dummy_balance = 236

	buttons = [
		[
			types.KeyboardButton("Find driver"),
			types.KeyboardButton("Find passenger"),
		],
		[
			types.KeyboardButton(f"Balance - {dummy_balance} Go"),
			types.KeyboardButton("History of trips"),
		],
		[
			types.KeyboardButton("Profile"),
		],
		[
			types.KeyboardButton("FAQ"),
			types.KeyboardButton("Support"),
		],
	]

	keyboard_markup = types.ReplyKeyboardMarkup(
		keyboard=buttons,
		resize_keyboard=True,
		input_field_placeholder="Select a menu item"
	)

	# TODO: normal database query for user name
	dummy_username = "Poul"

	await message.answer(text=f"{dummy_username}, have a nice trip!", reply_markup=keyboard_markup)


if __name__ == "__main__":
	config = dotenv_values(".env")

	bot = Bot(config['API_TOKEN'])
	dispatcher = Dispatcher(bot)

	dispatcher.register_message_handler(choose_language, commands=["start"])

	executor.start_polling(dispatcher, skip_updates=True)
