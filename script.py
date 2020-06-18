from flask import Flask, render_template, request, redirect, url_for

import requests
import json



app = Flask(__name__)


def getCountries():
    response = requests.get("http://localhost:11111/countries")
    data = response.json()
    countries = data['countries']
    return countries

def getId(name):
    countries = getCountries()
    for country in countries:
        if country['name'] == name:
            ident = country['id']
            break 
    ident = str(ident)
    return ident

@app.route("/")
def index():
    countries = getCountries()
    names = []
    for country in countries:
        names.append(country['name'])

    return render_template("index.html", names=names)

@app.route("/add")
def add():
    return render_template("add.html")

@app.route("/add", methods=['POST'])
def addCountry():
    name = request.form['name']
    pop = request.form['pop']
    cap = request.form['cap']
    dem = request.form['dem']
    flag_img = request.form['flag']
    map_img = request.form['map']
    if dem == 'yes':
        dem = True
    else: 
        dem = False
    pop= int(pop)
    data = {'name': name, 'population': pop, 'capital': cap, 'isDemocracy': dem, 'flag': flag_img, 'map': map_img }
    data = json.dumps(data)
    requests.post("http://localhost:11111/countries", data=data)
    countries = getCountries()
    names = []
    for country in countries:
        names.append(country['name'])
    return render_template("index.html", names=names)

@app.route("/remove/<name>")
def remove(name):
    ident = getId(name)
    requests.delete("http://localhost:11111/countries/" + ident)
    countries = getCountries()
    names = []
    for country in countries:
        names.append(country['name'])
    return render_template("index.html", names=names)

@app.route("/country/<name>")
def renderCountry(name):
    ident = getId(name)
    response = requests.get("http://localhost:11111/countries/" + ident)
    data =response.json()
    data = data['country']
    print(data)
    name = data['name']
    population = data['population']
    capital = data['capital']
    isDemocracy = data['isDemocracy']
    flag = data['flag']
    map_img = data['map']
    if isDemocracy:
        isDemocracy = "is"
    else:
        isDemocracy = "is not"
    return render_template("country.html", cap=capital, isDem=isDemocracy, pop=population, name=name, flag=flag, map=map_img)

def removeCountry():
    print("called")


if __name__ == '__main__':
  app.run(host='0.0.0.0', port=5000)