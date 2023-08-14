from django.http import HttpResponse
from django.shortcuts import render, redirect
from django.views.decorators.http import require_http_methods

from .models import Person

def index(request):
    persons = Person.objects.all()
    return render(request, "base.html", {"person_list": persons})

@require_http_methods(["POST"])
def add (request):
    name = request.POST["name"]
    company = request.POST["company"]
    notes = request.POST["notes"]
    if (name == ""):
        name = "NA"
    if (company == ""):
        company = "NA"
    if (notes == ""):
        notes = "NA"
    person = Person(name=name, company=company, notes=notes)
    person.save()
    return redirect("index")

def updatePage(request, person_id):
    person = Person.objects.get(id=person_id)
    return render(request, "update.html", {"person": person})

@require_http_methods(["POST"])
def update(request, person_id):
    person = Person.objects.get(id=person_id)
    
    name = request.POST["name"]
    company = request.POST["company"]
    notes = request.POST["notes"]
    if (name == ""):
        name = person.name
    if (company == ""):
        company = person.company
    if (notes == ""):
        notes = person.notes

    person.name = name
    person.company = company
    person.notes = notes
    person.save()
    return redirect("index")

def delete(request, person_id):
    person = Person.objects.get(id=person_id)
    person.delete()
    return redirect("index")