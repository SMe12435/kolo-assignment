<h1 class="code-line" data-line-start=0 data-line-end=1 ><a id="Marvel_Universe_Backend_Assignment_0"></a>Marvel Universe Backend Assignment</h1>
<h5 class="code-line" data-line-start=1 data-line-end=2 ><a id="In_a_nutshell_uses_REPL_to_fetch_characters_using_Marvel_Comic_API_APIs_are_also_exposed_which_can_be_used_by_other_clients_1"></a>In a nutshell, uses REPL to fetch characters using Marvel Comic API. APIs are also exposed which can be used by other clients.</h5>
<p class="has-line-data" data-line-start="3" data-line-end="4"><a href="https://developer.marvel.com/"><img src="https://seeklogo.com/images/M/Marvel_Comics-logo-D489AEB9C1-seeklogo.com.png" alt="Marvel API Url"></a></p>
<h2 class="code-line" data-line-start=5 data-line-end=6 ><a id="How_to_Setup_5"></a>How to Setup</h2>
<p class="has-line-data" data-line-start="6" data-line-end="9">Signup on Marvel Developers and generate the API Key<br>
Install GoLang and Redis<br>
Setup .env.dev file with Redis and Marvel API Key</p>
<p class="has-line-data" data-line-start="10" data-line-end="11">References</p>
<p class="has-line-data" data-line-start="12" data-line-end="15">To setup GoLang : <a href="https://go.dev/doc/install">https://go.dev/doc/install</a><br>
To setup Redis : <a href="https://redis.io/docs/getting-started/installation/">https://redis.io/docs/getting-started/installation/</a><br>
To know more about Gin Framework : <a href="https://github.com/gin-gonic/gin">https://github.com/gin-gonic/gin</a></p>
<h2 class="code-line" data-line-start=16 data-line-end=17 ><a id="To_run_16"></a>To run</h2>
<p class="has-line-data" data-line-start="18" data-line-end="19">In terminal, run :</p>
<pre><code class="has-line-data" data-line-start="20" data-line-end="22" class="language-sh">go run *.go
</code></pre>
<p class="has-line-data" data-line-start="23" data-line-end="24"><img src="https://dl3.pushbulletusercontent.com/mycMKXPXeEDyU0UE02Wli6BxKgTjIUyB/Screenshot%20from%202022-08-30%2018-03-54.png" alt="Terminal Sample SS"></p>
<p class="has-line-data" data-line-start="25" data-line-end="26">For help, type help in the terminal and it will display all the available commands</p>
<ul>
<li class="has-line-data" data-line-start="27" data-line-end="28">NS : New Search. Searches the Marvel Directory of Characters</li>
<li class="has-line-data" data-line-start="28" data-line-end="29">N  : Fetches Next Available Page</li>
<li class="has-line-data" data-line-start="29" data-line-end="30">P : Fetches Previous Available Page</li>
<li class="has-line-data" data-line-start="30" data-line-end="31">clear : clears the screen</li>
<li class="has-line-data" data-line-start="31" data-line-end="32">exit : exits the program</li>
<li class="has-line-data" data-line-start="32" data-line-end="33">help : displays help menu</li>
</ul>
<h2 class="code-line" data-line-start=36 data-line-end=37 ><a id="License_36"></a>License</h2>
<p class="has-line-data" data-line-start="38" data-line-end="39">MIT</p>
