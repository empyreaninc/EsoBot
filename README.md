# EsoBot
Main  automation module for Sh*pe project
Run from main file in IDE. 
Bot will display CLI with site options, select END & Entry module. 
Will check VPNREADY endpoint before requesting headers & sending request to Sh*pe protected API to avoid hanging on reconnecting VPN. 
Requests set of headers from EsoLinker server locally (start linker first) before using the grabbed headers to allow request to pass API protection. 
Entry module logs in, obtains relevant account details & enters according to raffle ID(must set sizeID to active raffle ID for this to work, current ID in payload is inactive).
All stages of this are displayed via CLI & retries are hardcoded currently, reason for failed request displayed as part of error message. Requires account with all fields filled in CSV to enter. Will iterate over entire account list in CSV to completion.
Account generation also works using same CSV as entry (E*ND.csv).
