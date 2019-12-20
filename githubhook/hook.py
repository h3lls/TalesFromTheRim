from github_webhook import Webhook
from flask import Flask
import os

app = Flask(__name__)  # Standard Flask app
webhook = Webhook(app) # Defines '/postreceive' endpoint

@app.route("/")        # Standard Flask endpoint
def hello_world():
    return "Hello, World!"

@webhook.hook()        # Defines a handler for the 'push' event
def on_push(data):
    print("Got push with: {0}".format(data))

@webhook.hook("pull_request")
def on_pull_request(data):
    # print(data['repository'])
    if data['action'] != 'merged' and data['action'] != 'closed':
        branch = data['pull_request']['head']['ref']
        clone_url = data['repository']['ssh_url']
        print(os.popen("/TalesFromTheRim/githubhook/builddev.sh {0} {1}".format(clone_url, branch)).read())

@webhook.hook("release")
def on_release(data):
    if data['action'] == 'published':
        print("Building for prod release for: {0}".format(data))

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8900)
