from flask import render_template
from flask_cors import CORS
from connexion.resolver import RestyResolver 
import config

# app = config.connex_app
# CORS(app.app, resources={r"/swagger/*": {"origins": "*"}}) 
# app.add_api(config.basedir / "swagger.yml", resolver=RestyResolver("api"))
app = config.connex_app 
CORS(app.app) 
app.add_api(config.basedir / "swagger.yml", resolver=RestyResolver("api"))


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000, debug=True)
    

   






