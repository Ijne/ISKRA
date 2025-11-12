package timepad

import (
	"fmt"
	"iskra/shared/config"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	// params.Add("cities", "Москва")
	params.Add("limit", "10")

	fullURL := baseURL + "?" + params.Encode()
	// log.Println("Full url: " + fullURL)
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

func (t *TimepadPoller) GetFilteredEvents(filter EventsFilter) ([]Event, error) {
	baseURL := t.timepadUrl + "events.json"
	params := url.Values{}

	params.Add("starts_at_min", time.Now().Format("2006-01-02T15:04:05-0700"))
	params.Add("starts_at_max", time.Now().Add(time.Duration(30)*time.Hour*24).Format("2006-01-02T15:04:05-0700"))
	if len(filter.City) != 0 {
		params.Add("cities", filter.City)
	}
	if filter.Limit != 0 {
		params.Add("limit", strconv.Itoa(filter.Limit))
	}
	if filter.Skip != 0 {
		params.Add("skip", strconv.Itoa(filter.Skip))
	}

	catIds := make([]string, 0)
	// категории
	for _, cat := range filter.Categories {
		newCat, ok := Categories[cat]
		log.Printf("- %v\n", newCat)

		if !ok {
			catIds = append(catIds, strconv.Itoa(463))
		} else {
			catIds = append(catIds, strconv.Itoa(int(newCat.ID)))
		}
	}
	log.Printf("* %v\n", catIds)
	if len(catIds) != 0 {
		params.Add("category_ids", strings.Join(catIds, ","))
	}

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

func (t *TimepadPoller) GetEventByID(eventID int64) (Event, error) {
	fullURL := t.timepadUrl + "events/" + strconv.Itoa(int(eventID))

	fmt.Println("Full url: " + fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Event{}, err
	}

	req.Header.Add("Authorization", "Bearer "+t.token)
	resp, err := t.client.Do(req)
	if err != nil {
		return Event{}, fmt.Errorf("error in request to timepad: %w", err)
	}

	var response Event
	err = render.DecodeJSON(resp.Body, &response)
	if err != nil {
		return Event{}, fmt.Errorf("error while unmarshaling timepad json: %w\n%v", err, resp.Body)
	}

	return response, nil
}

type EventsFilter struct {
	City       string
	Categories []string
	Limit      int
	Skip       int
}

type Category struct {
	Name string
	ID   int64
}

var Categories = map[string]Category{
	"Концерты": Category{"Концерты", 460},

	"Кино": Category{"Кино", 374},

	"Выставки": {"Выставки", 458},

	"Театры": Category{"Театры", 459},

	"Фестивали": Category{"Другие события", 462},

	"Спортивные события": Category{"Спорт", 376},

	"Вечеринки": Category{"Вечеринки", 457},

	"Клубы": Category{"Другие развлечения", 463},

	"Рестораны": Category{"Еда", 456},

	"Кафе": Category{"Еда", 456},

	"Пикники": Category{"Экскурсии и путешествия", 461},

	"Походы": Category{"Экскурсии и путешествия", 461},

	"Мастер классы": Category{"Хобби и творчество", 524},

	"Лекции": Category{"Наука", 2465},

	"Йога сессии": Category{"Спорт", 376},

	"Танцы": Category{"Спорт", 376},

	"Настольные игры": Category{"Интеллектуальные игры", 2335},

	"Караоке": Category{"Другие развлечения", 463},

	"Боулинг": Category{"Спорт", 376},

	"Картинг": Category{"Спорт", 376},
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

Концерты → Концерты (id: 460)

Кино → Кино (id: 374)

Выставки → Выставки (id: 458)

Театры → Театры (id: 459)

Фестивали → Другие события (id: 462)

Спортивные события → Спорт (id: 376)

Вечеринки → Вечеринки (id: 457)

Клубы → Другие развлечения (id: 463)

Рестораны → Еда (id: 456)

Кафе → Еда (id: 456)

Пикники → Экскурсии и путешествия (id: 461)

Походы → Экскурсии и путешествия (id: 461)

Мастер-классы → Хобби и творчество (id: 524)

Лекции → Наука (id: 2465)

Йога-сессии → Спорт (id: 376)

Танцы → Спорт (id: 376)

Настольные игры → Интеллектуальные игры (id: 2335)

Караоке → Другие развлечения (id: 463)

Боулинг → Спорт (id: 376)

Картинг → Спорт (id: 376)
*/
