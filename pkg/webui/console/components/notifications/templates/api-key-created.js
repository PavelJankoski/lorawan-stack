// Copyright © 2023 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { defineMessages } from 'react-intl'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import { getEntity } from '../utils'

const m = defineMessages({
  title: 'A new API key has just been created for your {entityType}',
  greeting: 'Dear {receiverName},',
  body: 'A new API key has just been created for your {entityType} <code>{id}</code> on your network.',
  apikey: '<b>API Key ID:</b> <code>{apiKeyId}</code>',
  rights: 'Rights:',
  right: '<code>{right}</code>',
  closing: 'You can view and edit this API key <Link>here</Link>.',
  preview:
    'A new API key has just been created for your {entityType} {id} on your network. API Key ID: {apiKeyId}',
})

const ApiKeyCreatedPreview = ({ notificationData }) => {
  const { entity_ids, data } = notificationData
  const { id } = data

  return (
    <Message
      content={m.preview}
      values={{
        entityType: getEntity(entity_ids),
        id: entity_ids[`${getEntity(entity_ids)}_ids`][`${getEntity(entity_ids)}_id`],
        apiKeyId: id,
        linebreak: <br />,
      }}
    />
  )
}

ApiKeyCreatedPreview.propTypes = {
  notificationData: PropTypes.shape({
    data: PropTypes.shape({
      id: PropTypes.string.isRequired,
    }).isRequired,
    entity_ids: PropTypes.shape({}).isRequired,
  }).isRequired,
}

const ApiKeyCreatedTitle = ({ notificationData }) => {
  const { entity_ids } = notificationData
  return (
    <Message
      content={m.title}
      values={{
        entityType: getEntity(entity_ids),
      }}
    />
  )
}

ApiKeyCreatedTitle.propTypes = {
  notificationData: PropTypes.shape({
    entity_ids: PropTypes.shape({}).isRequired,
  }).isRequired,
}

const ApiKeyCreated = ({ reciever, notificationData }) => {
  const { entity_ids, data } = notificationData
  const { id, rights } = data

  return (
    <>
      <Message content={m.greeting} values={{ receiverName: reciever }} component="p" />
      <Message
        content={m.body}
        values={{
          entityType: getEntity(entity_ids),
          id: entity_ids[`${getEntity(entity_ids)}_ids`][`${getEntity(entity_ids)}_id`],
          code: msg => <code>{msg}</code>,
          b: msg => <b>{msg}</b>,
        }}
        component="p"
      />
      <Message
        component="p"
        content={m.apikey}
        values={{
          b: msg => <b>{msg}</b>,
          code: msg => <code>{msg}</code>,
          apiKeyId: id,
        }}
      />
      <p>
        <Message component="b" content={m.rights} />
        <ul>
          {rights.map(right => (
            <>
              <Message
                component="li"
                content={m.right}
                values={{
                  code: msg => <code>{msg}</code>,
                  lineBreak: <br />,
                  right,
                }}
              />
              <Message content={{ id: `enum:${right}` }} firstToUpper />
            </>
          ))}
        </ul>
      </p>
      <Message
        content={m.closing}
        values={{
          Link: msg => (
            <Link
              to={`/applications/${
                entity_ids[`${getEntity(entity_ids)}_ids`][`${getEntity(entity_ids)}_id`]
              }/api-keys`}
            >
              {msg}
            </Link>
          ),
        }}
      />
    </>
  )
}

ApiKeyCreated.propTypes = {
  notificationData: PropTypes.shape({
    data: PropTypes.shape({
      id: PropTypes.string.isRequired,
      rights: PropTypes.arrayOf(PropTypes.string).isRequired,
    }).isRequired,
    entity_ids: PropTypes.shape({}).isRequired,
  }).isRequired,
  reciever: PropTypes.string.isRequired,
}

ApiKeyCreated.Title = ApiKeyCreatedTitle
ApiKeyCreated.Preview = ApiKeyCreatedPreview

export default ApiKeyCreated
