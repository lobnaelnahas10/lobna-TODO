import unittest
from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.chrome.options import Options
import todo_page

class todopage(unittest.TestCase):
    def setUp(self):
        chrome_options=Options()
        chrome_options.add_experimental_option("detach",True)
        self.driver=webdriver.Chrome(ChromeDriverManager().install())
        self.driver.get("http://localhost:3001/")
        self.driver.maximize_window()
        print('SETUP')
    def TiTle_Todo_page(self):
        todo_title=todo_page.todo(self.driver)
        print(self.driver.title)
        assert todo_title.is_the_title_correct()
    def test_Navigate_Create_task(self):
        todo_page.todo(self.driver).Create_New_task()

    def test_Navigate_Delete_task(self):
        todo_page.todo(self.driver).Delete_task_By_Mark()
    
    def test_Navigate_Update(self):
        todo_page.todo(self.driver).Update_task()
    
    def test_Navigate_Get_task_ByID(self):
        todo_page.todo(self.driver).Get_task_By_ID()

    def test_CheckBox_Mark(self):
        todo_page.todo(self.driver).Checkbox_Mark()

    def test_task_withNotExsiting_ID(self):
        todo_page.todo(self.driver).Get_task_By_NoID()

    def test_Delete_task_With_UnexsitingID(self):
        todo_page.todo(self.driver).Delete_task_By_NoID()

    def test_Createtask_Without_taskname(self):
        todo_page.todo(self.driver).Create_New_task_without_taskname()

    def test_Createtask_withCharID(self):
        todo_page.todo(self.driver).Create_New_task_with_CharID()

    def test_Createtask_ExsitingID(self):
        todo_page.todo(self.driver).Create_New_task_WithExsitingID()

    def tearDown(self):
        self.driver.quit()
        print('DONE')


if __name__ == "__main__":
    unittest.main()