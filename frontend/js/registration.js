document.getElementById('registrationPage').innerHTML = `
    <div id="re">
    <h2>Register</h2>
    <form id="registrationForm" action="/register" method="POST">
    <div class="flex">
        <label for="FirstName">First Name:</label>
        <input class="ui1" type="text" id="First Name" name="FirstName"><br>

        <label for="LastName">Last Name:</label>
        <input class="ui1" type="text" id="LastName" name="LastName"><br>
</div>
        <label for="nickname">Nickname:</label>
        <input class="ui1" type="text" id="regNickname" name="nickname"><br>

          <label for="age">Age:</label>
        <input class="ui1" type="number" id="regAge" name="Age"><br>

        <label for="email">Email:</label>
        <input class="ui1" type="email" id="email" name="email"><br>

        <label for="password">Password:</label>
        <input class="ui1" type="password" id="regPassword" name="password"><br>

         <label for="Gender">Gender:</label>
        <select class="ui1" id="Gender" name="Gender">
        <option value="Male">Male</option>
        <option value="female">Female</option>
    </select><br><br>

        <button type="submit">Register</button>
    </form>

    </div>
`;


if(formName){
    formName.addEventListener("submit", (e) => {
        e.preventDefult();
        
        
    })
}