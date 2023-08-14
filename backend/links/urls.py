from django.urls import path

from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("add", views.add, name="add"),
    path("delete/<int:person_id>/", views.delete, name="delete"),
    path("updatePage/<int:person_id>/", views.updatePage, name="updatePage"),
    path("update/<int:person_id>/", views.update, name="update"),
]