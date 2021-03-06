 
# MotorHome Club Specification

## Overview

The motorhome club is a community of Motorhome owners, and this web application provides a number of useful functions that benefit Motorhome users and club members. The club is built upon the Freemium model, ensuring that it is easy for new entry members to join up. Additional features and functions are made available to full members.
The app provides a sites directory including club sites( Motorhome suitable campsite ) Pub Stops, Wild Camping locations and StopOvers ( Other facilities such as attractions, museums, local craft producers e.g cheese makers candle makers etc )
The app provides journey / activity planning, with a mapping feature that enables you to draw a circle / rectangle on a map area, and provide a list of things you're interested in, and it provides pictures and locations pulled from flickr.
The app provides a GeoCaching function, that allows members to perform national treasure hunts. 
The app is tightly integrated with Facebook and the Motohome club Facebook page, which acts as the central point for social interaction ( why re-invent the wheel )
the club sells memberships, promotional marketing ( careful with this one GDPR ) it sells reviews, interviews and video production, along with editorial placement in the club magazine, and the YouTube channel and Podcast

# Feature Specifications

## Users

Users can join our free tier membership registering with email, and password or using their Social Media account. There are 4 users types

 * Guest ( free tier )
 * Member ( Standard paid membership )
 * Founder ( Premium members )
 * Patrons ( Commercial memberships )
 
ALL users must register, and very very limited access is provide to web hoy poloy! Membership is £48 per year. Founder memberships are earned rather than bought ( criteria to be established ). Patron membership is ( maybe £250 per year, or £125 if they offer our members a discount )

Users have a profile including a username, email, birthday ( year of birth not required ), location ( usually city or town ), a profile picture

### Users data table

id, username, first_name, last_name, email, birthday, location, images_id, member_type (guest, member, founder, patron - relates to lookup table) joined_date, last_seen_date, logon_count, login_type ( username, google, facebook, twitter, etc..) membership_expires

The data provides some useful metrics about logons when last seen to help with marketing and community promotion

Implement last_seen_date and logon_count via an activity_audit table. This holds logon, logoff, timeout date time details for every user. last_seen_date is then implemented as a func on the activity_audit model, and is logon_count.

## Sites

Sites are locations that where club members can stay in their motorhomes, there are 4 types club_sites, pub_stops, wild_camping, and stop_overs. Sites are owned, and maintained by Patron members. A site record provides Edit functionality ONLY to the user registered as the sites owner. Patron members have the ability to create club_site listings. Each site listing has an annual listing fee of £24. Updates and amendments are free, listings expire on their anniversary and requiring renewal with a listing fee. 

### Sites data table

id, name, description, attractions, email, telephone, website, facilities( use a referencing table, so you can have differences at each site ), sitetype, owninguser ( related user_id ), images(images which are to appear in the site carousel, are recorded in the images table and referenced via the images.related_id), longitude, latitude, addressstreet, locality, town, county, postcode, listingexpires( Date when listing should no longer be shown) 

After some thinking, and in regard to managing access to Sites for various member types ( which presented a problem in that, I could not see a way of passing the resource being access to the authorisor)
I have decided to implement sites as 4 discrete resources ( campsites, pustops, stopovers, and wildcamps) This enables me to wrap the resources with the appropriate authorisation.
This also enables me to group resources by authorisation type. 
However, this is not very SOLID in that we're creating more resources than we need, and more authorizers than we need. But, this does fit well in the sense that whenever a user tries to access a resource
that their membership does not provide, they get a message telling that, rather than an empty data set.
However, I have kept the  data structure (as described above ) the same for each model, so that this can be refactored out later if required.
## Images function and data table

All site images are held in the images table, and stored in the ./images/<image_type>/ directory on the filesystem (SysAdmin:- Backups will be required), there are 3 image_type directories

 * /originals
 * /resized 
 * /thumbnails 
 
id, related_id, original, resized, thumbnail

Images are processed, saved and retrieved by the imageManager function. The function provides to interfaces

IF1 - Accepts a map of images ( which have been uploaded ) it saves Originals, creates a resized and cropped version, and a thumbnail of the resized image, it writes these to file, and creates an entry for them in the images data table with a unique id, and a related_id for the batch. Thus, an upload of 8 images, would generate 8 rows in the database table, with unique id's, but all 8 rows would hold the same related_id, and return the related_id for the related record to store. 

IF2 - Accepts a related_id ( held by the related record), and returns a map of image URL's

## Club Caches

Caches provide a geo_cache style activity. Caches are placed in interesting locations, that are worthy of a visit by members. Caches can be placed on any of our sites. for club_sites and pub_stops the physical site owner must give permission for the cache to be placed, and must be willing for club members to visit without staying or incurring a fee to retrieve the cache. Caches have a unique "Playing Card"(secretnumber, secretname e.g Ace Clubs) and the Cache Found register uses this identifier to determine whether the registering users is making a legitimate claim of the find.
Cache placements are also reviewing and featured in the Club Magazine ( No Fixed Abode )

Caches are validated once they are first found by someone other than the deployer. Caches are maintained via a status (Deployed, Found, Muggled, Lost, Missing, Removed)
A notification mechanism is employed that updates the Cache seeder, when ever someone finds the cache. A confirmation mechanism is provided within each cache, where by a playing card is taped to this inside of the cache tub lid. A finder registering the find must correctly state the Suit, and the Face ( 1,2,3...J,Q,K etc..), this confirms their find, and logs it against their Cache Log. It also sends an email to the Seeder, letting them know that the user found their cache. Club Caches implement the Clap and Comment functions.

### Club Cache data model

cachename description:null.text secretnumber secretname images longitude latitude validated:false.bool deployedby onsite status

buffalo g resource caches cachename description:null.text secretnumber secretname images longitude latitude validated:false.bool deployedby onsite status 

## Events 

This is to be developed as a later feature, but will offer something similar to the club_sites where event organisors can become Patrons, and register their events for a fee.

## Payment provision

# Other Functions

## Claps

As per Medium Claps allow users to rate site and cache records by using the Claps feature. A user may clap a record upto 10 times, effectively giving it a rating from 1 to 10. Claps are stored independently of the record, keeping a record of how many claps from each user has been applied to the site or cache record.

## Mag-e-zine

TODO - Based upon the Medium model, but also provides a monthly generated PDF, which is generated from content created that month... TODO

# Membership Benefits, and Monetisation

I'm not sure whether I like the idea of paid for membership levels, as that is not very inclusive. I notice that Free Code Camp use Flash Card Pop ups to try to encourage their uses to take out a monthly subscription.
I was thinking that this could work well for MotorHome Club. It needs some form of badgeing system, perhaps with something applied to the users avatar, that lets other users know.

Perhaps levels could be Guest ( Black n White avatar ), Nomad ( £2.00 per month, £20 per year ), Pilgrim ( £4.00, £40 per year ) , and a Founder ( £6.00, £ 60 per year ) - Yearly subscribers get 2 months for free.
Founder members also get the benefit of being placed on the site reviewers register, which grants them free stays at sites across the UK whilst performing a site or attraction review.