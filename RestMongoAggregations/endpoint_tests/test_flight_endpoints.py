import requests, os, pytest


def pytest_configure():
    pytest.TOKEN = ""
    pytest.first_flight_id = ""


def test_init():
    req = requests.get("http://localhost:8080/init-flight")
    assert req.status_code == 200
    req = requests.get("http://localhost:8080/init-user")
    assert req.status_code == 200


def test_login():
    req = requests.post(
        url="http://localhost:8080/login",
        json={"username": "admin1", "password": "123"},
    )
    pytest.TOKEN = req.headers["Bearer"]
    assert pytest.TOKEN.startswith("ey")


def test_create_flight():
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
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.status_code == 201


def test_get_all_flights():
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
        {"Belgrade", 99, 100, 111, "Subotica", "2025-01-01T00:00:00Z"},
        {"Belgrade", "Novi Sad", 112, 90, "2025-01-01T00:00:00Z"},
        {58, 200, "London", "Budapest", 150, "2023-05-06T12:00:42.123Z"},
    ]


def test_get_flight_by_id():
    req = requests.get(
        url="http://localhost:8080/get-flight-by-id?id=" + pytest.first_flight_id
    )
    assert req.json()[0] == {
        "id": pytest.first_flight_id,
        "date": "2025-01-01T00:00:00Z",
        "endPlace": "Belgrade",
        "startPlace": "Subotica",
        "capacity": 100,
        "price": 111,
        "freeSeats": 99,
    }


def test_update_flight():
    req = requests.patch(
        url="http://localhost:8080/flight/" + pytest.first_flight_id,
        json={
            "id": pytest.first_flight_id,
            "capacity": 200,
            "date": "2023-12-31T12:00:42.123Z",
            "endPlace": "London",
            "freeSeats": 50,
            "price": 150,
            "startPlace": "Budapest",
        },
    )
    req = requests.get(
        url="http://localhost:8080/get-flight-by-id?id=" + pytest.first_flight_id
    )
    assert req.json()[0] == {
        "id": pytest.first_flight_id,
        "capacity": 200,
        "date": "2023-12-31T12:00:42.123Z",
        "endPlace": "London",
        "freeSeats": 50,
        "price": 150,
        "startPlace": "Budapest",
    }


def test_delete_flight():
    req = requests.delete(
        url="http://localhost:8080/flight/" + pytest.first_flight_id,
        json={
            "id": pytest.first_flight_id,
            "capacity": 200,
            "date": "2023-12-31T12:00:42.123Z",
            "endPlace": "London",
            "freeSeats": 50,
            "price": 150,
            "startPlace": "Budapest",
        },
    )

    req = requests.get(url="http://localhost:8080/flight")

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
        {"Novi Sad", "Belgrade", 112, "2025-01-01T00:00:00Z", 90},
        {58, 200, "Budapest", 150, "London", "2023-05-06T12:00:42.123Z"},
    ]


# def test_create_flight_two():
#     req = requests.post(
#         url="http://localhost:8080/flight",
#         json={
#             "date": "2023-05-06T12:00:42.123Z",
#             "endPlace": "London",
#             "startPlace": "Budapest",
#             "capacity": 200,
#             "price": 150,
#             "freeSeats": 58,
#         },
#         headers={"Bearer": pytest.TOKEN},
#     )
#     assert req.status_code == 201


def test_create_flights():
    req = requests.post(
        url="http://localhost:8080/flight",
        json={
            "date": "2023-12-31T14:00:42.123Z",
            "endPlace": "Belgrade",
            "startPlace": "Amsterdam",
            "capacity": 100,
            "price": 50,
            "freeSeats": 100,
        },
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.status_code == 201

    req = requests.post(
        url="http://localhost:8080/flight",
        json={
            "date": "2023-12-31T14:00:42.123Z",
            "endPlace": "Belgrade",
            "startPlace": "Maldivi",
            "capacity": 100,
            "price": 50,
            "freeSeats": 100,
        },
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.status_code == 201

    req = requests.post(
        url="http://localhost:8080/flight",
        json={
            "date": "2023-12-30T15:00:42.123Z",
            "endPlace": "Belgrade",
            "startPlace": "Maldivi",
            "capacity": 100,
            "price": 50,
            "freeSeats": 100,
        },
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.status_code == 201

    req = requests.post(
        url="http://localhost:8080/flight",
        json={
            "date": "2023-12-30T11:00:42.123Z",
            "endPlace": "Budapest",
            "startPlace": "Istanbul",
            "capacity": 100,
            "price": 50,
            "freeSeats": 100,
        },
        headers={"Bearer": pytest.TOKEN},
    )
    assert req.status_code == 201
