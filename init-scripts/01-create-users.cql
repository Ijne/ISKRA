WITH range(1, 1000) AS ids
UNWIND ids AS id

WITH id,
     ['IT and Technology', 'Design and UX', 'Medicine', 'Education', 'Business', 'Finance', 'Marketing', 'Art', 'Music', 'Cinema', 'Photography', 'Architecture', 'Engineering', 'Real Estate', 'Jurisprudence', 'Psychology'] AS careers,
     ['Extrovert', 'Introvert', 'Ambivert', 'Analyst', 'Creator', 'Pragmatist', 'Romantic', 'Realist', 'Optimist', 'Philosopher', 'Innovator', 'Leader', 'Purposeful', 'Flexible', 'Persistent', 'Team Player'] AS personalities,
     ['Serious Relationships', 'Friendship', 'Casual Relationships', 'Family Creation', 'Partner Search', 'Romance', 'Business Partnership', 'Creativity', 'Travel', 'Joint Projects', 'Spirituality', 'Career'] AS relationships,
     ['Love and Care', 'Family', 'Career', 'Finance', 'Spirituality', 'Health', 'Education', 'Creativity', 'Freedom', 'Adventure', 'Security', 'Ecology'] AS values_list

CREATE (p:Person {
    id: id,
    name: 'User_' + id + '_t',
    username: '@user' + id + '_t',
    age: 18 + (id % 50),
    gender: id % 2,
    preferred_gender: id % 3,
    career_type: careers[(id * 7) % size(careers)],
    personality_type: personalities[(id * 11) % size(personalities)],
    relationship_goal: relationships[(id * 13) % size(relationships)],
    important_values: values_list[(id * 17) % size(values_list)],
    currentLimit: 30,
    lastLimitUpdate: toString(datetime()),
    maxLimit: 30,
    refreshAmount: 5,
    refreshRate: "PT1H",
    music: '',
    films: '',
    hobbies: '',
    event_preferences: '',
    photo: ''
});