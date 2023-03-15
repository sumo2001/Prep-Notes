- https://docs.google.com/document/d/1PiHcXLJ93zOmeY-1BToRr-X0Az4vt7UWMN1rQBw7luI/edit?usp=sharing
- Requirement Total number of SSH failed login attempts: <b>source="secure.log" action=failure </b>  
- Find how many failed logins from every IPs: <b>source="secure.log" action=failure | stats count by src</b>  
- Find List of Countries from which the failed login attempts were made: <b>source="secure.log" action=failure | iplocation src </b>
- Create a visualization of countries in world map based on failed logins: <b>source="secure.log" action=failure | iplocation src | geostats count by Country</b>
- The Splunk search processing language (SPL) supports the Boolean operators : AND, OR, and NOT.
- We can automate reports generated from splunk to deliver them to a email, for that we need to configure a SMTP server in splunk server settings
- Splunk APP or mostly used for dashboards
- https://docs.google.com/document/d/137z48GMwp4okkvAZhxUD-TXaS9z93N-FaqyUq7IB6uQ/edit?usp=sharing

