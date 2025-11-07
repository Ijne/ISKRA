package timepad

import (
	"fmt"
	"iskra/shared/config"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/render"
)

// var orgs = []string{"abc"}

type TimepadPoller struct {
	client     *http.Client
	token      string
	timepadUrl string
}

func New(cfg *config.Config) *TimepadPoller {
	poller := TimepadPoller{}
	poller.client = &http.Client{}
	poller.token = cfg.Timepad.Token
	poller.timepadUrl = "https://api.timepad.ru/v1/"

	return &poller
}

func (t *TimepadPoller) GetEvents() ([]Event, error) {
	baseURL := t.timepadUrl + "events.json"
	params := url.Values{}

	params.Add("starts_at_min", time.Now().Format("2006-01-02T15:04:05-0700"))
	params.Add("starts_at_max", time.Now().Add(time.Duration(30)*time.Hour*24).Format("2006-01-02T15:04:05-0700"))
	params.Add("cities", "Москва")
	params.Add("limit", "10")

	fullURL := baseURL + "?" + params.Encode()
	log.Println("Full url: " + fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+t.token)
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in request to timepad: %w", err)
	}

	var response Response
	err = render.DecodeJSON(resp.Body, &response)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling timepad json: %w\n%v", err, resp.Body)
	}

	return response.Values, nil
}

/*
все категории событий:
{
  "values": [
    {
      "id": 217,
      "name": "Бизнес",
      "tag": "business"
    },
    {
      "id": 374,
      "name": "Кино",
      "tag": "cinema"
    },
    {
      "id": 376,
      "name": "Спорт",
      "tag": "sport"
    },
    {
      "id": 379,
      "name": "Для детей",
      "tag": "kids"
    },
    {
      "id": 382,
      "name": "Иностранные языки",
      "tag": "languages"
    },
    {
      "id": 399,
      "name": "Красота и здоровье",
      "tag": "beauty"
    },
    {
      "id": 452,
      "name": "ИТ и интернет",
      "tag": "it"
    },
    {
      "id": 453,
      "name": "Психология и самопознание",
      "tag": "psychology"
    },
    {
      "id": 456,
      "name": "Еда",
      "tag": "food"
    },
    {
      "id": 457,
      "name": "Вечеринки",
      "tag": "party"
    },
    {
      "id": 458,
      "name": "Выставки",
      "tag": "exhibition"
    },
    {
      "id": 459,
      "name": "Театры",
      "tag": "theater"
    },
    {
      "id": 460,
      "name": "Концерты",
      "tag": "concert"
    },
    {
      "id": 461,
      "name": "Экскурсии и путешествия",
      "tag": "trip"
    },
    {
      "id": 462,
      "name": "Другие события",
      "tag": "other_event"
    },
    {
      "id": 463,
      "name": "Другие развлечения",
      "tag": "other_entertainment"
    },
    {
      "id": 524,
      "name": "Хобби и творчество",
      "tag": "hobby"
    },
    {
      "id": 525,
      "name": "Искусство и культура",
      "tag": "art"
    },
    {
      "id": 1315,
      "name": "Образование за рубежом",
      "tag": "education_abroad"
    },
    {
      "id": 1940,
      "name": "Гражданские проекты",
      "tag": "civil"
    },
    {
      "id": 2335,
      "name": "Интеллектуальные игры",
      "tag": "intellekt"
    },
    {
      "id": 2465,
      "name": "Наука",
      "tag": "science"
    }
  ]
}
*/
