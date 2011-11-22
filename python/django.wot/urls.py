# URL layout:
# subverse.info/wiki
# subverse.info/who
# subverse.info/




from django.conf.urls.defaults import patterns, include, url

# Uncomment the next two lines to enable the admin:
from django.contrib import admin
admin.autodiscover()

urlpatterns = patterns('',
    # Examples:
    # url(r'^$', 'subverse1.views.home', name='home'),
    # url(r'^subverse1/', include('subverse1.foo.urls')),
#      url(r'^about/', 'subverse1.views.about', name='about'),
    # Uncomment the admin/doc line below to enable admin documentation:
     url(r'^admin/doc/', include('django.contrib.admindocs.urls')),

    # Uncomment the next line to enable the admin:
     url(r'^admin/', include(admin.site.urls)),
)
