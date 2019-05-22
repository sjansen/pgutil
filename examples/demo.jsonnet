local decrypt(target) = {
  [target + '/decrypted']: {
    target: target + '/echo',
    after: [target + '/reverse', target + '/rotate'],
  },
  [target + '/reverse']: {
    target: target + '/rev',
  },
  [target + '/rotate']: {
    target: target + '/rot13',
  },
};

local encrypted(msg) = {
  class: 'demo',
  config: {
    string: msg,
  },
};

{
  targets: {
    msg1: encrypted('!abbcF'),
    msg2: encrypted('.rvy n fv rxnp ruG'),
    msg3: encrypted('.revczni n fv qyebj ruG'),
  },
  tasks: decrypt('msg1')
         + decrypt('msg2')
         + decrypt('msg3'),
}
