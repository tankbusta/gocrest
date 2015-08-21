# gocrest

A golang library for communicating with EVE Online's Public CREST API.
 
You probably shouldn't use this yet, but feel free to contribute!

## Notes/Assumptions

* EVE Online's API is strange that it includes the same information twice although sometimes varying types (string, integer, etc). Gocrest will only return the correct type based on it's value. If you need a string representation of say the total # of alliances... convert it yourself. 
* EVE's API also returns href's in some of the response objects that point to a more detailed version of the data in question. Instead of including this during unmarshalling, a `GetHref()` function will be provided.

# License

This library may contain static data which is owned by CCP. Otherwise gocrest is licensed under the MIT License.

© 2014 CCP hf. All rights reserved. "EVE", "EVE Online", "CCP", and all related logos and images are trademarks or registered trademarks of CCP hf.