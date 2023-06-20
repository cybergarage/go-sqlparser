CREATE TABLE usertable (
            YCSB_KEY TEXT,
            FIELD0 TEXT, FIELD1 TEXT, FIELD2 TEXT, FIELD3 TEXT,
            FIELD4 TEXT, FIELD5 TEXT, FIELD6 TEXT, FIELD7 TEXT,
            FIELD8 TEXT, FIELD9 TEXT,
            PRIMARY KEY (YCSB_KEY ASC))
            SPLIT AT VALUES (('user10'),('user14'),('user18'),
            ('user22'),('user26'),('user30'),('user34'),('user38'),
            ('user42'),('user46'),('user50'),('user54'),('user58'),
            ('user62'),('user66'),('user70'),('user74'),('user78'),
            ('user82'),('user86'),('user90'),('user94'),('user98'));