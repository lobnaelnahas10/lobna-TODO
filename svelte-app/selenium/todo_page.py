from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By

from locators import *

class BasePage(object):
    def __init__(self,driver):
        self.driver=driver

class todo(BasePage):
    #Navigate title
    def is_the_title_correct(self):
        return 'TODO' in self.driver.title
    #Navigate create new task
    def Create_New_task(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('10')
        self.driver.find_element(By.XPATH,todopage.Task).send_keys('task10')
        self.driver.find_element(By.XPATH,todopage.CreateBtn).click()
    #Navigate Delete task by delete mark
    def Delete_task_By_Mark(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.DeleteMark))).click()
    #Navigate Delete task by ID
    def Delete_task_By_ID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('9')
        self.driver.find_element(By.XPATH,todopage.DeleteBtn).click()
    #Navigate Update task 
    def Update_task(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('9')
        self.driver.find_element(By.XPATH,todopage.Task).send_keys('task900')
        self.driver.find_element(By.XPATH,todopage.UpdateBtn).click()
    #Navigate Get task by ID
    def Get_task_By_ID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('999')
        self.driver.find_element(By.XPATH,todopage.GetByIDBtn).click()
    #Navigate Checkbox task by delete mark
    def Checkbox_Mark(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.CheckBoxMark))).click()
    #Getby Not exsiting ID
    def Get_task_By_NoID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('25')
        self.driver.find_element(By.XPATH,todopage.GetByIDBtn).click()
    #Delete task by not exsiting ID
    def Delete_task_By_NoID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('30')
        self.driver.find_element(By.XPATH,todopage.DeleteBtn).click()
    #Create without enter task name
    def Create_New_task_without_taskname(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('88')
        self.driver.find_element(By.XPATH,todopage.CreateBtn).click()
    #Create with Char ID
    def Create_New_task_with_CharID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('Amira')
        self.driver.find_element(By.XPATH,todopage.Task).send_keys('projec')
        self.driver.find_element(By.XPATH,todopage.CreateBtn).click()
    #Create with Exsiting ID
    def Create_New_task_WithExsitingID(self):
        wait=WebDriverWait(self.driver,60)
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.ID))).send_keys('2')
        wait.until(EC.visibility_of_element_located((By.XPATH,todopage.Task))).send_keys('Hisham')
        self.driver.find_element(By.XPATH,todopage.CreateBtn).click()   