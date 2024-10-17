document.getElementById('loginPage').innerHTML = `
    <div id="re">
    <h2>Login</h2>
    <form id="loginForm" action="/login" method="POST">
        <label for="email">Email:</label>
        <input class="ui1" type="email" id="email" name="email"><br>
        
        <label for="password">Password:</label>
        <input class="ui1" type="password" id="password" name="password"><br>
        
        <button type="submit">Login</button>
    </form>
    </div>
`;
