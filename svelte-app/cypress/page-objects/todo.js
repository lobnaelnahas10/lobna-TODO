class Todopage{
    getFromtitle()
    {
        return cy.get('h2.svelte-ke1wqu')//.contains
    }
    getIDbtn()
    {
        return cy.get('#inputID')   
    }
    gettaskbtn()
    {
        return cy.get('#inputTask')
    }
    getcreatebtn()
    {
        return cy.get('[type="create"]')
    }
    getDeletebtn()
    {
        return cy.get('[type="delete"]')
    }
    getupdatebtn()
    {
        return cy.get('.txt_filed').find('[type="update"]')
    }
    getGetByIDbtn()
    {
        return cy.get('[type="get"]')
    }
    getDeletekmark()
    {
        return cy.get(':nth-child(4) > .btn > i')
    }
    getCheckmark()
    {
        return cy.get(':nth-child(6) > input')
    }

    //cases of existing ID
    //show title of page
    showtitle()
    {
        this.getFromtitle().contains("TODO")
    }
    //check delete button
    deleteTodo()
    {
        this.getIDbtn().type("1")
        this.getDeletebtn().click()
    }
    //check create button
    //put id and task name then click on button
    createTodo()
    {
        this.getIDbtn().type("30")
        this.gettaskbtn().type("proj30")
        this.getcreatebtn().click()
    }
    //check get button
    //put id then click on button
    getTodo()
    {
        this.getIDbtn().clear({force: true}).type("9")
        this.getGetByIDbtn().click()
    }
    //check update button
    //put id then new task name then click on button
    updateTodo()
    {
        this.getIDbtn().type("999")
        this.gettaskbtn().type("proj11")
        this.getupdatebtn().click()
    }
    //check checkbox of mark
    //click on it then make it empty then click which task you want to check
    checkmarkTodo()
    {
        this.getFromtitle().contains("TODO")
        this.getCheckmark().click()
        this.getIDbtn().type(" ")
        this.getIDbtn().type(2)
        this.getGetByIDbtn()
    }
    deleteTodoBymark()
    {
        this.getDeletekmark().click()
    }

    //check if i delete task with unexsiting id
    //first clear then type id
    deleteTodo_UnExsiting_ID()
    {
        this.getIDbtn().clear({force: true}).type("55")
        this.getDeletebtn().click()
    }
    //check if it will create task without task name
    //first clear then type id only
    //then click on create button
    createTodo_without_task()
    {
        this.getIDbtn().clear({force: true}).type("60")
        this.getcreatebtn().click()
    }
    //check if it will create with char id
    //first clear then type char id and task name
    //then click on create button
    createTodo_withCharID()
    {
        this.getIDbtn().clear({force: true}).type("Lobna")
        this.gettaskbtn().type("test")
        this.getcreatebtn().click()
    }
    //check if it will create with exsiting id
    //first clear then type char id and task name
    //then click on create button
    createTodo_withExsiting_ID()
    {
        this.getIDbtn().clear({force: true}).type("9")
        this.gettaskbtn().type("Task9")
        this.getcreatebtn().click()
    }
    
}
export default new Todopage();