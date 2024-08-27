# Generated by Django 3.2.13 on 2022-07-12 08:03

from django.db import migrations, models
from django.db.models import F
import django_migration_linter as linter


def fill_rotation_start_field(apps, schema_editor):
    CustomOnCallShift = apps.get_model("schedules", "CustomOnCallShift")
    CustomOnCallShift.objects.update(rotation_start=F("start"))


class Migration(migrations.Migration):

    dependencies = [
        ('schedules', '0005_auto_20220704_1947'),
    ]

    operations = [
        linter.IgnoreMigration(),
        migrations.AddField(
            model_name='customoncallshift',
            name='rotation_start',
            field=models.DateTimeField(default=None, null=True),
        ),
        migrations.RunPython(fill_rotation_start_field, migrations.RunPython.noop),
        migrations.AlterField(
            model_name='customoncallshift',
            name='rotation_start',
            field=models.DateTimeField(),
        ),
    ]