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

local message(data) = {
  class: 'strbuf',
  config: {
    data: data,
  },
};

{
  targets: {
    msg1: message('!abbcF'),
    msg2: message('.rvy n fv rxnp ruG'),
    msg3: message('.revczni n fv qyebj ruG'),
  },
  tasks: decrypt('msg1')
         + decrypt('msg2')
         + decrypt('msg3'),
}
