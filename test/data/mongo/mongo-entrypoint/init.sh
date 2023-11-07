#! /bin/bash
echo "Creating users..."
mongosh admin --host localhost -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --eval "db.createUser(
    {
        user: '$MONGO_IM_USERNAME', 
        pwd: '$MONGO_IM_PASSWORD',
        roles: [
            {
                role: 'readWrite', 
                db: 'mongo'
            }
        ]
    });"
echo "Users created"
