from flask import Flask, request
from telegram_bots.logger import setup_logger

app = Flask(__name__)

logger = setup_logger()

@app.route('/healthcheck', methods=['GET'])
def backend_failure() -> tuple:
    logger.info('Healthcheck endpoint was hit with a healthcheck call.')
    return {'status': 'success'}, 200
