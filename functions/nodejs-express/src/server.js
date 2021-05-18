const express = require("express");
const bodyParser = require("body-parser");

const app = express();

const PORT = process.env.PORT || 3000;

app.use(bodyParser.json());

// Action Handler
// See: https://hasura.io/docs/latest/graphql/core/actions/action-handlers.html#action-handlers

app.post('/hi', async (req, res) => {
  const ua = req.header('User-Agent') || '-';

  // access control
  const secret = req.header('ACTION_SECRET_DUMMY') || '-';
  if (secret !== "dummy" ) {
    return res.status(400).json({
      message: "Do NOT access!",
      code: "FORBIDDEN",
    });
  }


  const successResponse = {
    hello: "hasura",
    userAgent: ua,
    message: `secret:${secret}` || '-',
    //foo: "foo",
  };
  const errorResponse = {
    message: "Something wrong happened:(",
    code: "ERROR_TEST",
  };
  const random = Math.floor(Math.random() * 2);
  if (random === 1) {
    return res.json(successResponse);
  } else {
    return res.status(400).json(errorResponse);
  }
});

app.post('/hello', async (req, res) => {
  const ua = req.header('User-Agent') || '-';
  const { message } = req.body.input;
  return res.json({
    hello: "world",
    userAgent: ua,
    message: message || '-',
    // foo: "foo",
  });
});

// Event Triggers
// https://hasura.io/docs/latest/graphql/core/event-triggers/payload.html
app.post('/createTodo', async (req, res) => {
  const ua = req.header('User-Agent') || '-';
  const payload = req.body;
  console.log(`id:${payload.id}`);
  console.log(`createdAt:${payload.created_at}`);
  console.log(`trigger.name:${payload.trigger.name}`);
  console.dir(`event:${JSON.stringify(payload.event)}`);

  res.status(204).send();
});


app.listen(PORT);
