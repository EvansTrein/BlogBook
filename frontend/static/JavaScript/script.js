fetch('http://localhost:7000')
  .then(response => response.json())
  .then(data => {
    console.log(data);
    const dataDiv = document.getElementById('data');
    dataDiv.innerHTML = data.message;
  })
  .catch(error => console.error('Ошибка:', error));