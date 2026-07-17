from django.contrib import admin

from .models import Profile

# Register your models here.
# we have to register the models here for to be able to see in admin dashboard of django
admin.site.register(Profile)
