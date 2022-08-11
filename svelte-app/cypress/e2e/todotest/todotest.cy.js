import todo from"../../page-objects/todo"
describe('CypressTest', function(){

    before(function(){
        cy.visit('/')
    })

    it('Create New task',function(){
        todo.createTodo()
    })

    it('Delete task by button',function(){
        todo.deleteTodo()
    })
    it('Delete task by mark',function(){
        todo.deleteTodoBymark()
    })

    it('Update task',function(){
        todo.updateTodo()
    })

    it('Get task By ID',function(){
        todo.getTodo()
    })
    it('Check task by mark',function(){
        todo.checkmarkTodo()
    })
    it('Show title',function(){
        todo.showtitle()
    })
    //check of unexsiing key
    it('Delete task by unexsiting ID ',function(){
        todo.deleteTodo_UnExsiting_ID()
    })
    it('Create task without input task ',function(){
        todo.createTodo_without_task()
    })

    it('Create task with char ID ',function(){
        todo.createTodo_withCharID()
    })
    it('Create task with Exsiting ID ',function(){
        todo.createTodo_withExsiting_ID()
    })
    after(function(){
        cy.log('Done')
    })
})