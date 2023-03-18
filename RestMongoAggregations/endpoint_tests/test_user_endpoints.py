import requests, os, jwt, pytest

def pytest_configure():
    pytest.TOKEN = ""
    
def test_init():
    req = requests.get("http://localhost:8080/init")
    req = requests.get("http://localhost:8080")
    assert req.status_code == 200 

def test_create_user():
    req = requests.post(url="http://localhost:8080", json={"username": "naz1", "password": "123"})
    req = requests.get("http://localhost:8080/")
    #checking only username and passoword, not ID
    assert list(map(lambda x: {x["username"], x["password"]},req.json())) == [{"naz1", "123"}]