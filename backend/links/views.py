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
    
    person = Person(name=name, company=company, notes=notes)
    person.save()
    return redirect("index")

def update(request, person_id):
    person = Person.objects.get(id=person_id)
    person.company = "Microsoft"
    person.save()
    return redirect("index")

def delete(request, person_id):
    person = Person.objects.get(id=person_id)
    person.delete()
    return redirect("index")