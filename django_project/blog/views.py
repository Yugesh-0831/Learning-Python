from django.http import HttpResponse
from django.shortcuts import render

from .models import Post

posts = [
    {
        "author": "Yugesh",
        "title": "Blog Post 1",
        "content": "First post content",
        "date_posted": "Feb 7 2026",
    },
    {
        "author": "Aman",
        "title": "Blog Post 2",
        "content": "Second post content",
        "date_posted": "Feb 7 2026",
    },
]


# Create your views here.
def home(request):
    context = {"posts": Post.objects.all()}
    return render(request, "blog/home.html", context)


def about(request):
    return render(request, "blog/about.html")
