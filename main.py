from dotenv import dotenv_values
from aiogram import Bot, Dispatcher, executor

if __name__ == "__main__":
	config = dotenv_values(".env")

	bot = Bot(config['API_TOKEN'])
	dispatcher = Dispatcher(bot)
