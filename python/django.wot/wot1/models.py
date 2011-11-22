from django.db import models

# Create your models here.
class User(models.Model):
    keyid = models.CharField(max_length=40)
    username = models.CharField(max_length=40)
    def __unicode__(self):
        return self.keyid + " " + self.username


class Rating(models.Model):
    boolR = models.NullBooleanField("Boolean rating", blank=True, null=True)
    intR = models.IntegerField("Integer rating", blank=True, null=True)
    comment1= models.CharField(max_length=180, blank=True, null=True)
    def __unicode__(self):
        return "Bool: "+self.boolR.__str__()+", Int: "+self.intR.__str__()

class Link(models.Model):
    source = models.ForeignKey(User, related_name='sent_links')
    sink = models.ForeignKey(User, related_name='rx_links')
    rating = models.ForeignKey(Rating)
    def __unicode__(self):
        return (self.source.keyid.__str__() + " -> " + self.sink.keyid.__str__() + " : "
                                            + self.rating.__str__())
