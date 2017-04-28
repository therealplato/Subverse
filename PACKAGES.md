Packages provide rulescapes and interfaces describing a particular set of behaviors

comms:
  - note to self (1 party)
  - whisper (2)
  - chat (n)
  - publish (∞)

identity:
  - make keys
hygiene:
  transmission policies:
  - keep to self
  - share only to <Identity...>
  - limited share, n hops
  - limited share, n copies
  - limited share, t ms
  - unlimited share: ∞ hops

  retention policies:
  - burn after reading
  - burn after t ms
  - burn after sharing
  - at will
  - keep forever

  modification policies:
  - burn data only
  - burn metadata only
  - append your metadata
  - replace metadata

content:
  Ref() Hash
  metadata:
  - SignWork(MetaContent, Identity) MetaContent
  - SetHygienePolicy(MetaContent, HygienePolicy)
  policies:
  - share 
  - share with metadata appended
  usage:
  - fork
  - remix with
  subvert:
  - copy internet content into subverse
  index:
  - ContentAvailable(Hash) bool