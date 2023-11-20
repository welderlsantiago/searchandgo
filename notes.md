<h1>Notas Go</h1>

<h2>Useful Commands</h2>
<p>go mod init (git repository link) -> intialize Go modules in your project</p>

<p>go tidy -> manage imports of your project, what is being used in you project will be added to your module and what is not will be removed</p>

<h2>Organization</h2>

<ul>
<li>Separate routes by handler methods (SHOW, CREATE, DELETE, PUT...)</li>
<li>Create a specific package folder for the schema</li>
<li>Create a specifir package folder for the Database, frequently called 'config'</li>
</ul>