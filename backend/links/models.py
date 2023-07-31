from django.db import models

# Create your models here.
class Person(models.Model):
    name=models.CharField(max_length=350)
    company=models.CharField(max_length=350)
    notes=models.CharField(max_length=1000)
    
    def __str__(self):
        return self.name