
# Antimiasma Vaccination Report

Some repositories on your machine have been infected with the Miasma worm.
This vaccine used the same spreading mechanism to be able to beat the worm.

More details at https://github.com/Team-Rockstars-Security/antimiasma

### Infected Repositories

The list of infected repositories on your system was:

{MESSAGES}

### IMPORTANT

Check all those repositories for used credentials, tokens, passwords,
AND ssh keys. Rotate and change all of those, including company-wide tokens,
they've been compromised by the Miasma worm.

The Antimiasma implant contains the EICAR test string, so if you have an IT
security department that uses a SIEM/SOAR system to track their EDR alerts,
they will automatically have been notified. Please make sure to tell them
what happened and give them a copy of this file.

The Antimiasma program will remove itself after a timeout of 7 days since the
first infection.

In the future, please make sure to not use AI assisted IDEs in non-sandboxed
repositories anymore, as this was the cause for the infection. The autostart
integrations of AI assisted IDEs are an attack surface that pretty much no
EDR is able to keep track of.

Have a great day!

