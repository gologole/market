const participantTemplate = `
<p>Имя</p>
<input type="text" placeholder="" name="participant-name">
<p>Специфика</p>
<select name="participant-options">
    <option>Frontend</option>
    <option>Backend</option>
    <option>Дизайнер</option>
</select>
<button class="search-button">Поиск</button>
`;

const teamTemplate = `
<p>Название команды</p>
<input type="text" placeholder="" name="team-name">
<div class=buttons-wrapper>
    <button class="search-button">Поиск</button>
    <button class="search-button">Создать команду</button>
</div>
`;
const participantsList = document.querySelector('.participants-list');
const fetchAdress = 'http://localhost:8080'

function updateForm() {
    const formFields = document.getElementById('form-fields');
    const roleSelect = document.getElementById('role-select');
    const h2 = document.getElementById('h2')
    const selectedRole = roleSelect.value;

    formFields.innerHTML = '';

    if (selectedRole === 'participant') {
        formFields.innerHTML = participantTemplate;
        h2.textContent = 'Участники';
        fetchUsers();
    } else if (selectedRole === 'team') {
        formFields.innerHTML = teamTemplate;
        h2.textContent = 'Команды';
        participantsList.innerHTML = '';
    }
}

// Инициализировать форму и пользователей при загрузке страницы
document.addEventListener('DOMContentLoaded', () => {
    const roleSelect = document.getElementById('role-select');
    roleSelect.value = 'participant';
    updateForm();
    fetchUsers();
}); 

function fetchUsers() {
    fetch(`${fetchAdress}/profiles`)
        .then(response => response.json())
        .then(data => {
            displayUsers(data);
        })
        .catch(error => console.error('Error fetching data:', error));
}

function displayUsers(users) {
    participantsList.innerHTML = '';

    users.forEach(user => {
        const listItem = document.createElement('li');
        listItem.className = 'participant';
        listItem.onclick = () => openPopup(user);


        const img = document.createElement('img');
        img.className = 'user-image-small';
        img.src = user.page;
        img.alt = `Profile Picture of ${user.name}`;

        const name = document.createElement('p');
        name.className = 'user-name';
        name.innerText = `${user.name}`;

        listItem.appendChild(img);
        listItem.appendChild(name);
        participantsList.appendChild(listItem);
    });
}

function openPopup(user) {
    document.getElementById('popup-name').innerText = `${user.first_name} ${user.last_name}`;
    document.getElementById('popup-image').src = user.avatar;
    document.getElementById('popup-description').innerText = `Email: ${user.email}`;    
    document.getElementById('popup').style.display = 'flex';
}

function closePopup() {
    document.getElementById('popup').style.display = 'none';
}
