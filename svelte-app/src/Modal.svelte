<script>
    import { onMount } from "svelte/internal";
    
    let todos = []
    const baseURL = "http://localhost:8080/todo"
    
    const getTodos = async () => {
      const response = await fetch(baseURL)
      const data = await response.json()
      todos = await data
    }
    
    onMount(() => {
      getTodos()
    })
    
    
    let Task
    let ID
    const createTodo = async event => {
      event.preventDefault() 
      await fetch(baseURL, {
        method: "post",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: ID,
          task: Task,
        }),
      })
    
      getTodos()
      ID=""
      Task = ""
    }
    const deleteTodo = async (todo) => {
      event.preventDefault()
      await fetch(baseURL+"/"+todo.id,{
        method: "delete",
      })
      getTodos()
    }
  
    const deleteTodobyID= async (todo) => {
      event.preventDefault()
      await fetch(baseURL+"/"+ID,{
        method: "delete",
       
      })
     
      getTodos()
    }
    var returned= null
    const GetbyId= async (todo) => {
      event.preventDefault()
      const res=await fetch(baseURL+"/"+ID,{
        method: "get",
        headers: {
          "Content-Type": "application/json",
          Accept:"application/json"
        },
  
       
      })
      if (!res.ok) {
        throw new Error(`Error! status: ${res.status}`);
      }
  
      // 👇️ const result: CreateUserResponse
      const result = (await res.json()) ;
      returned=  JSON.stringify(result, null, 4)
      console.log('result is: ', JSON.stringify(result, null, 4));
  
      getTodos()
      
      
    }
  
    const UpdateByID= async (todo) => {
      event.preventDefault()
      await fetch(baseURL+"/"+todo.ID,{
        method: "put",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: ID,
          task: Task,
        }),
       
      })
     
      getTodos()
    }
    const markTask = async (todo) => {
      event.preventDefault()
      await fetch(baseURL + "/"+todo.id,{
        method: "patch",
      })
      getTodos()
    }
    
    
    </script>
    
    <main>
      <h2>TODO List</h2>
    <div class="center">
      <form on:submit="{createTodo}">
        <div class="txt_filed">
          <input type="text" bind:value="{ID}" id="inputID">
          <input type="text" bind:value="{Task}" id="inputTask">
          <button type="create"  on:click="{createTodo}"> create </button>
          <button type="delete" on:click="{deleteTodobyID}" > delete </button>
          <button type="update" on:click="{UpdateByID}" > Update </button>
          <button type="get" on:click="{GetbyId}" > GetID </button>
          <div>{returned} </div>
      </form>
      <hr  style="color: #BC3CBC" >
      <hr  style="color: #BC3CBC" >
    
      {#each todos as todo}
      <div>
        <input type="checkbox" onclick={(e) => markTask(todo)}>
        <h3>{todo.id} {todo.task}</h3>
        <button name="btn" class="btn" on:click={(e) => deleteTodo(todo)}><i calss="fa fa-trash">❌</i></button>
      </div>
      {/each}
    </div>
    </main>
    
      
    
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <style>
  
     .center{
      position: absolute;
      top: 70%;
      left: 50%;
      transform: translate(-50%,-50%);
      width: 400px;
      background: white;
      border-radius: 10px ;
  
    }
   
    
      input[type=text]{
        margin: 8px 0;
        border-radius: 20px;
    
      }
    
      #inputID{
        width: 70%;
        box-sizing: border-box;
        padding: 12px 20px;
        border: 1px solid rgb(20, 64, 209);
      }
    
      #inputTask{
        width: 70%;
        box-sizing: border-box;
        padding: 12px 20px;
        border: 1px solid rgb(20, 64, 209);
      }
    
      button[type=create]{
        background-color: rgb(39, 11, 117);
        border: none;
        color: rgb(244, 244, 248);
        padding: 20px 32px;
        font: size 30%;;
        text-decoration: none;
        margin: 6px 6px;
        cursor: pointer;
        border-radius: 25px;
      }
       
      button[type=delete]{
        background-color: rgb(39, 11, 117);
        border: none;
        color: rgb(244, 244, 248);
        padding: 20px 32px;
        font: size 30%;;
        text-decoration: none;
        margin: 4px 4px;
        cursor: pointer;
        border-radius: 25px;
      }
      button[type=update]{
        background-color: rgb(39, 11, 117);
        border: none;
        color: rgb(244, 244, 248);
        padding: 20px 32px;
        font: size 30%;;
        text-decoration: none;
        margin: 4px 4px;
        cursor: pointer;
        border-radius: 25px;
      }
      button[type=get]{
        background-color: rgb(39, 11, 117);
        border: none;
        color: rgb(244, 244, 248);
        padding: 20px 32px;
        font: size 30%;;
        text-decoration: none;
        margin: 4px 4px;
        cursor: pointer;
        border-radius: 25px;
      }
    
      h2{
        color: rgb(214, 219, 220);
        font-size: 35px;
        font-weight: 400;
        background-image: linear-gradient(to left, #fdfdff, #d1bce5);
      }
    
      .btn {
  
      background-color: white; 
      border: none; 
      color: white; 
      padding: 12px 16px; 
      font-size: 16px; 
      position: absolute;
    }
    
    button, h3{
      display: inline-block;
    }
    
    .btn:hover {
      background-color: rgb(255, 255, 255);
    }
    
    h3{
      font-weight: bold;
      color: rgb(0, 0, 128);
      padding-right: 33px;
      font-style: italic;
      margin-left: 9px;
      
    }
    h2 {
    color: #000000;
    text-transform: uppercase;
    font-size: 3em;
    font-weight: 100;
  }
    </style>