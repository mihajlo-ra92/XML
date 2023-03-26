import requests, pytest

def pytest_configure():
    pytest.TOKEN = ""
    pytest.first_user_id = ""
    pytest.first_flight_id = ""

def test_init():
    req = requests.get("http://localhost:8080/init-ticket")
    assert req.status_code == 200
    req = requests.get("http://localhost:8080/init")
    assert req.status_code == 200
    req = requests.get("http://localhost:8080/init-flight")
    assert req.status_code == 200


def test_create_user():
    req = requests.post(
        url="http://localhost:8080/user",
        json={
            "username": "nazTicket",
            "password": "123",
            "userType": "regular",
            "email": "nazTicket@gmail.com",
            "firstName": "Fname2",
            "lastName": "Lname2",
            "birthDate": "1609891200",
            "gender": "male",
            "governmentId": "333444555",
        },
    )
    assert req.status_code == 201

def test_login():
    req = requests.post(
        url="http://localhost:8080/login", json={"username": "nazTicket", "password": "123"}
    )
    pytest.TOKEN = req.json()["Bearer"]
    print(pytest.TOKEN)
    assert pytest.TOKEN.startswith("ey")
def test_check_user():
    resp = requests.get("http://localhost:8080/user", headers={"Bearer": pytest.TOKEN})
    pytest.first_user_id = resp.json()[0]["id"]
    assert list(
        map(lambda x: {x["username"], x["password"], x["userType"], x["birthDate"]}, resp.json())
    ) == [{"naz1", "123", "regular", "2025-01-01T00:00:00Z"}, {"nazTicket", "123", "regular", "2021-01-06T00:00:00Z"}]
def test_create_flight_two():
    req = requests.post(
        url="http://localhost:8080/flight",
        json={
            "date": "2023-05-06T12:00:42.123Z",
            "endPlace": "London",
            "startPlace": "Budapest",
            "capacity": 200,
            "price": 150,
            "freeSeats": 58,
        },
    )
    assert req.status_code == 201
    req = requests.get(url="http://localhost:8080/flight")
    pytest.first_flight_id = req.json()[0]["id"]
    # checking everythig but ID
    assert list(
        map(
            lambda x: {
                x["date"],
                x["endPlace"],
                x["startPlace"],
                x["capacity"],
                x["price"],
                x["freeSeats"],
            },
            req.json(),
        )
    ) == [{"2023-05-06T12:00:42.123Z", "London", "Budapest", 200, 150, 58}]

 

def test_create_ticket():
    req = requests.post(
        url="http://localhost:8080/ticket", headers={"Bearer": pytest.TOKEN}, json={}
    )