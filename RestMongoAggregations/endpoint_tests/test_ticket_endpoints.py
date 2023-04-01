import requests, pytest


def pytest_configure():
    pytest.TOKEN = ""
    pytest.first_user_id = ""
    pytest.first_flight_id = ""


def test_init():
    req = requests.get("http://localhost:8080/init-ticket")
    assert req.status_code == 200
    req = requests.get("http://localhost:8080/init-user")
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
        url="http://localhost:8080/login",
        json={"username": "nazTicket", "password": "123"},
    )
    pytest.TOKEN = req.json()["Bearer"]
    print(pytest.TOKEN)
    assert pytest.TOKEN.startswith("ey")


def test_check_user():
    resp = requests.get("http://localhost:8080/user", headers={"Bearer": pytest.TOKEN})
    pytest.first_user_id = resp.json()[3]["id"]
    assert list(
        map(
            lambda x: {x["username"], x["password"], x["userType"], x["birthDate"]},
            resp.json(),
        )
    ) == [
        {"2025-01-01T00:00:00Z", "123", "admin", "admin1"},
        {"us1", "123", "regular", "2025-01-01T00:00:00Z"},
        {"2025-01-01T00:00:00Z", "us2", "123", "regular"},
        {"123", "regular", "nazTicket", "2021-01-06T00:00:00Z"},
    ]


def test_create_flight_two():
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
    ) == [
        {"Subotica", 99, 100, 111, "Belgrade", "2025-01-01T00:00:00Z"},
        {112, "Belgrade", 90, "2025-01-01T00:00:00Z", "Novi Sad"},
    ]


def test_create_ticket():
    req = requests.post(
        url="http://localhost:8080/ticket",
        headers={"Bearer": pytest.TOKEN},
        json={
            "date": "2023-05-06T12:00:42.123Z",
            "endPlace": "London",
            "startPlace": "Budapest",
            "capacity": 1,
            "price": 150,
            "flightId": pytest.first_flight_id,
            "userId": pytest.first_user_id,
        },
    )
    assert True


def test_get_tickets_by_user():
    req = requests.get(
        url="http://localhost:8080/get-tickets-by-user-id",
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.json() == [
        {
            "id": req.json()[0]["id"],
            "date": "2023-05-06T12:00:42.123Z",
            "endPlace": "London",
            "startPlace": "Budapest",
            "capacity": 1,
            "price": 150,
            "flightId": pytest.first_flight_id,
            "userId": pytest.first_user_id,
        }
    ]
