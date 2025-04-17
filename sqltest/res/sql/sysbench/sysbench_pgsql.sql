CREATE TABLE sbtest (
  id SERIAL,
  k INTEGER DEFA ULT '0' NOT NULL ,
  c CHAR(120)  DEFAULT '' NOT NULL,
  pad CHAR( 60) DEFAULT '' NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO sbtest1(k, c, pad) VALUES(5, '83868641912-28773972837-60736120486-75162659906-27563526494-20381887404-41576422241-93426793964-56405065102-33518432330', '67847967377-48000963322-62604785301-91415491898-96926520291'),(6, '38014276128-25250245652-62722561801-27818678124-24890218270-18312424692-92565570600-36243745486-21199862476-38576014630', '23183251411-36241541236-31706421314-92007079971-60663066966'),(6, '33973744704-80540844748-72700647445-87330233173-87249600839-07301471459-22846777364-58808996678-64607045326-48799346817', '38615512647-91458489257-90681424432-95014675832-60408598704'),(6, '37002370280-58842166667-00026392672-77506866252-09658311935-56926959306-83464667271-94685475868-28264244556-14550208498', '63947013338-98809887124-59806726763-79831528812-45582457048'),(5, '44257470806-17967007152-32809666989-26174672567-29883439075-95767161284-94957565003-35708767253-53935174705-16168070783', '34551750492-67990399350-81179284955-79299808058-21257255869'),(6, '37216201353-39109531021-11197415756-87798784755-02463049870-83329763120-57551308766-61100580113-80090253566-30971527105', '05161542529-00085727016-35134775864-52531204064-98744439797'),(6, '33071042495-29920376648-91343430102-79082003121-73317691963-02846712788-88069761578-14885283975-44409837760-90760298045', '91798303270-64988107984-08161247972-12116454627-22996445111'),(5, '73754818686-04889373966-18668178968-56957589012-31352882173-91882653509-59577900152-88962682169-52981807259-62646890059', '76460662325-41613089656-42706083314-81833284991-17063140920'),(5, '26482547570-00155460224-12388481921-23289186371-78242522654-77998886134-73270876420-50821093220-31442690639-11588920653', '30508501104-50823269125-88107014550-70202920684-95842308929'),(6, '05677017559-47107518969-97509137401-28934334557-14497052050-61906823704-44077628507-24840441785-05187301456-27797851637', '29489382504-13697582598-09964978366-26554639515-36136545002');

ALTER TABLE sbtest1 ADD INDEX k_1 (k);

SELECT c FROM sbtest1 WHERE id=$1;