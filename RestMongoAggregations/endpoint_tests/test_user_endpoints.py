import requests, os, pytest

def pytest_configure():
    pytest.TOKEN = ""
    pytest.first_user_id = ""
    
def test_init():
    req = requests.get("http://localhost:8080/init")
    assert req.status_code == 200 

def test_login():
    req = requests.post(url="http://localhost:8080/login", json={"username": "naz1", "password": "123"})
    pytest.TOKEN = req.json()['Bearer']
    print(pytest.TOKEN)
    assert pytest.TOKEN.startswith("ey")

def test_create_user_invalid():
    req = requests.post(url="http://localhost:8080/user", json={"username": "naz2", "password": "123", "userType":"regular"})
    assert req.status_code == 400 

def test_create_user():
    req = requests.post(url="http://localhost:8080/user", json={"username": "naz2", "password": "123", "userType":"regular", "email":"naz2@gmail.com", "firstName": "Fname2", "lastName":"Lname2"})
    assert req.status_code == 201
    resp = requests.get("http://localhost:8080/user", headers={"Bearer":pytest.TOKEN})
    pytest.first_user_id = resp.json()[0]['id']
    #checking only username and passoword, not ID
    assert list(map(lambda x: {x["username"], x["password"], x["userType"]},resp.json())) == [{"naz1", "123", "admin"}, {"naz2", "123", "regular"}]

def test_read_by_username():
    req = requests.get("http://localhost:8080/user/read-by-username?username=naz1")
    assert list(map(lambda x: {x["username"], x["password"], x["userType"]},req.json())) == [{"naz1", "123", "admin"}]

def test_patch_user():
    req = requests.patch("http://localhost:8080/user/"+pytest.first_user_id, json={"username": "naz1_update", "password": "123_update"})
    req = requests.get("http://localhost:8080/user/read-by-username?username=naz1_update")
    assert list(map(lambda x: {x["username"], x["password"], x["userType"]},req.json())) == [{"naz1_update", "123_update", "admin"}]


def test_delete_user():
    _ = requests.delete("http://localhost:8080/user/"+pytest.first_user_id)
    resp = requests.get("http://localhost:8080/user", headers={"Bearer":pytest.TOKEN})
    assert list(map(lambda x: {x["username"], x["password"], x["userType"]},resp.json())) == [{"naz2", "123", "regular"}]