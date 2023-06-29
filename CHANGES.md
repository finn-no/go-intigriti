# Intigriti V2

## Pre-existing functions

GetSubmissions(programId string) ---changed to---> GetProgramSubmissions(programId string)

## Functions added by Radu Boian

**GetSubmissions** -> gets all submissions
**GetSubmission** -> gets all submission by code
**GetProgramSubmissions** -> gets all submissions by program id
**GetSubmissionPayouts** -> get submission payouts by submission code
**GetSubmissionEvents** -> get submission events by submission code

## Structs added by Radu Boian

- Program
- Attachment
- Payout
- Event
- User

## Changed structs by Radu Boian

Changed Submission struct by refactoring Report.Attachments into a standalone structure Attachments.
Changed Submission struct by adding the fields Payouts and Events by the custom structures created (this helps for better indexing into Elasticsearch and searching through events)
Created User struct with all needed/possible fields. When GetSubmission is called, the User object in the JSON is filled only with userId, userName, role. It misses the email and other fields.
Added the following fields to subtype From and To in Event: Status, CloseReason, DuplicateSubmissionURL, UserID, UserName, AvatarURL, Email, Role

### Issues

Attachment structure is missing URL field
