# Generated by Django 3.2.18 on 2023-05-24 03:54

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('user_management', '0011_auto_20230411_1358'),
        ('alerts', '0015_auto_20230508_1641'),
        ('base', '0003_delete_organizationlogrecord'),
        ('twilioapp', '0003_auto_20230408_0711'),
    ]

    state_operations = [
        migrations.CreateModel(
            name='SMSRecord',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('exceeded_limit', models.BooleanField(default=None, null=True)),
                ('grafana_cloud_notification', models.BooleanField(default=False)),
                ('created_at', models.DateTimeField(auto_now_add=True)),
                ('status', models.PositiveSmallIntegerField(blank=True, choices=[(10, 'accepted'), (20, 'queued'), (30, 'sending'), (40, 'sent'), (50, 'failed'), (60, 'delivered'), (70, 'undelivered'), (80, 'receiving'), (90, 'received'), (100, 'read')], null=True)),
                ('sid', models.CharField(blank=True, max_length=50)),
                ('notification_policy', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='base.usernotificationpolicy')),
                ('receiver', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.CASCADE, to='user_management.user')),
                ('represents_alert', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='alerts.alert')),
                ('represents_alert_group', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='alerts.alertgroup')),
            ],
            options={
                'db_table': 'twilioapp_smsmessage',
            },
        ),
        migrations.CreateModel(
            name='PhoneCallRecord',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('exceeded_limit', models.BooleanField(default=None, null=True)),
                ('created_at', models.DateTimeField(auto_now_add=True)),
                ('grafana_cloud_notification', models.BooleanField(default=False)),
                ('status', models.PositiveSmallIntegerField(blank=True, choices=[(10, 'queued'), (20, 'ringing'), (30, 'in-progress'), (40, 'completed'), (50, 'busy'), (60, 'failed'), (70, 'no-answer'), (80, 'canceled')], null=True)),
                ('sid', models.CharField(blank=True, max_length=50)),
                ('notification_policy', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='base.usernotificationpolicy')),
                ('receiver', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.CASCADE, to='user_management.user')),
                ('represents_alert', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='alerts.alert')),
                ('represents_alert_group', models.ForeignKey(default=None, null=True, on_delete=django.db.models.deletion.SET_NULL, to='alerts.alertgroup')),
            ],
            options={
                'db_table': 'twilioapp_phonecall',
            },
        ),
    ]

    operations = [
        migrations.SeparateDatabaseAndState(state_operations=state_operations)
    ]

