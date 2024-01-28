# DiabolGo

Inspired by https://thehackernews.com/2024/01/china-backed-hackers-hijack-software.html to build a proof of concept demonstrating the attack technique.

This repo serves as an interception suite proof of concept capable of delivering malware without ever needing touch the target PC (Windows) as an attacker. No vulnerability scans, no software exploits involving memory overflows. I'll demonstrate in a video how as an APT you are able to craft and develop a highly targeted attack on your victim, simply by monitoring HTTP traffic. This could be a cool suite of other tools for intercept based attacks in the future.

As this is only a POC this is not designed to operate on the network layer of a router, this will not handle TLS traffic and is only designed to intercept HTTP for proof of concept. To fully weaponise this, one would have to take the concept which is demonstrated here, and build a module for a target router / switch / firewall on the network layer, with the same capabilities. 

I might build that in the future, but that's far more time consuming than building a POC.

# Legal notice

This project, including all associated source code and documentation, is developed and shared solely for educational, research, and defensive purposes in the field of cybersecurity. It is intended to be used exclusively by cybersecurity professionals, researchers, and educators to enhance understanding, develop defensive strategies, and improve security postures.

Under no circumstances shall this project be used for criminal, unethical, or any other unauthorized activities. The re-engineering and analysis of the malware sample provided in this project are meant to serve as a resource for learning and should not be employed for offensive operations or actions that infringe upon any individual's or organization's rights or privacy.

The author of this project disclaim any responsibility for misuse or illegal application of the material provided herein. By accessing, studying, or using this project, you acknowledge and agree to use the information contained within strictly for lawful purposes and in a manner that is consistent with ethical guidelines and applicable laws and regulations.

USE AT YOUR OWN RISK. If you decide to use this software CONDUCT A THOROUGH INDEPENDENT CODE REVIEW to ensure it meets your standards. No unofficial third party dependencies are included to minimise attack surface of a supply chain risk. I cannot be held responsible for any problems that arise as a result of executing this, the burden is on the user of the software to validate its safety & integrity. All care has been taken to write safe code.

It is the user's responsibility to comply with all relevant local, state, national, and international laws and regulations related to cybersecurity and the use of such tools and information. If you are unsure about the legal implications of using or studying the material provided in this project, please consult with a legal professional before proceeding. Remember, responsible and ethical behaviour is paramount in cybersecurity research and practice. The knowledge and tools shared in this project are provided in good faith to contribute positively to the cybersecurity community, and I trust they will be used with the utmost integrity.