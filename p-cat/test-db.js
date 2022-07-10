const { mongoose } = require('mongoose');

const Schema = mongoose.Schema;

// connect DB
mongoose.connect('monogdb://localhost/test-db', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});
// create schema
const PhotoSchema = new Schema({
  title: String,
  description: String,
});

const Photo = mongoose.model('Photo', PhotoSchema);
// create a photo
Photo.create({
  title: 'Photo title',
  description: 'loarem ipsum',
});
// read a photo
Photo.find({}, (err, data) => {
  console.log(data);
});
// update a photo
const id = '6079f04e5916c524d4bdcb74';
Photo.findByIdAndUpdate(
  id,
  {
    title: 'Photo title update',
    description: 'loarem ipsum update',
  },
  {
    new: true,
  },
  (err, data) => {
    console.log(data);
  }
);
// delete a photo
const id = '6079f04e5916c524d4bdcb74';
Photo.findByIdAndDelete(id, (err, data) => {
  console.log('Fhoto is removed..');
});
