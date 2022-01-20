package ad

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/microsoft/kiota/abstractions/go/serialization"

	"github.com/amanenk/cq-provider-msgraph/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

func AdGroups() *schema.Table {
	return &schema.Table{
		Name:        "msgraph_ad_groups",
		Description: "Group",
		Resolver:    fetchAdGroups,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"tenant_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "tenant_id",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeString,
				Resolver:    client.ResolveTenantId,
			},
			{
				Name:        "entity_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				//Resolver:    schema.PathResolver("DirectoryObject.Entity.additionalData"),
				Resolver: func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
					p, ok := resource.Item.(graph.Group)
					if !ok {
						return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
					}
					data := p.GetAdditionalData()
					resource.Set(c.Name, data)
				},
			},
			{
				Name:        "entity_id",
				Description: "Read-only.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DirectoryObject.Entity.id"),
			},
			{
				Name:     "deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DirectoryObject.deletedDateTime"),
			},
			{
				Name:        "allow_external_senders",
				Description: "Indicates if people external to the organization can send messages to the group",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("allowExternalSenders"),
			},
			{
				Name:        "auto_subscribe_new_members",
				Description: "Indicates if new members added to the group will be auto-subscribed to receive email notifications",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("autoSubscribeNewMembers"),
			},
			{
				Name:        "calendar",
				Description: "The group's calendar",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsCalendar,
			},
			{
				Name:        "calendar_view",
				Description: "The calendar view for the calendar",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsCalendarView,
			},
			{
				Name:        "classification",
				Description: "Describes a classification for the group (such as low, medium or high business impact)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("classification"),
			},
			{
				Name:        "conversations",
				Description: "The group's conversations.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsConversations,
			},
			{
				Name:        "created_date_time",
				Description: "Timestamp of when the group was created",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("createdDateTime"),
			},
			{
				Name:        "created_on_behalf_of_entity_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("createdOnBehalfOf.Entity.additionalData"),
			},
			{
				Name:        "created_on_behalf_of_id",
				Description: "Read-only.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("createdOnBehalfOf.Entity.id"),
			},
			{
				Name:     "created_on_behalf_of_deleted_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("createdOnBehalfOf.deletedDateTime"),
			},
			{
				Name:        "description",
				Description: "An optional description for the group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("description"),
			},
			{
				Name:        "display_name",
				Description: "The display name for the group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("displayName"),
			},
			{
				Name:        "drive",
				Description: "The group's default drive",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsDrive,
			},
			{
				Name:        "drives",
				Description: "The group's drives",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsDrives,
			},
			{
				Name:        "events",
				Description: "The group's calendar events.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsEvents,
			},
			{
				Name:        "expiration_date_time",
				Description: "Timestamp of when the group is set to expire",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("expirationDateTime"),
			},
			{
				Name:        "extensions",
				Description: "The collection of open extensions defined for the group",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsExtensions,
			},
			{
				Name:        "group_types",
				Description: "Specifies the group type and its membership",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("groupTypes"),
			},
			{
				Name:        "has_members_with_license_errors",
				Description: "Indicates whether there are members in this group that have license errors from its group-based license assignment",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("hasMembersWithLicenseErrors"),
			},
			{
				Name:        "hide_from_address_lists",
				Description: "True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("hideFromAddressLists"),
			},
			{
				Name:        "hide_from_outlook_clients",
				Description: "True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("hideFromOutlookClients"),
			},
			{
				Name:     "is_archived",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("isArchived"),
			},
			{
				Name:        "is_assignable_to_role",
				Description: "Indicates whether this group can be assigned to an Azure Active Directory role or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("isAssignableToRole"),
			},
			{
				Name:        "is_subscribed_by_mail",
				Description: "Indicates whether the signed-in user is subscribed to receive email conversations",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("isSubscribedByMail"),
			},
			{
				Name:        "license_processing_state_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("licenseProcessingState.additionalData"),
			},
			{
				Name:     "license_processing_state_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("licenseProcessingState.state"),
			},
			{
				Name:        "mail",
				Description: "The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("mail"),
			},
			{
				Name:        "mail_enabled",
				Description: "Specifies whether the group is mail-enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("mailEnabled"),
			},
			{
				Name:        "mail_nickname",
				Description: "The mail alias for the group, unique for Microsoft 365 groups in the organization",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("mailNickname"),
			},
			{
				Name:        "membership_rule",
				Description: "The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("membershipRule"),
			},
			{
				Name:        "membership_rule_processing_state",
				Description: "Indicates whether the dynamic membership processing is on or paused",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("membershipRuleProcessingState"),
			},
			{
				Name:        "onenote",
				Description: "Read-only.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsOnenote,
			},
			{
				Name:        "on_premises_domain_name",
				Description: "Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("onPremisesDomainName"),
			},
			{
				Name:        "on_premises_last_sync_date_time",
				Description: "Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("onPremisesLastSyncDateTime"),
			},
			{
				Name:        "on_premises_net_bios_name",
				Description: "Contains the on-premises netBios name synchronized from the on-premises directory",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("onPremisesNetBiosName"),
			},
			{
				Name:        "on_premises_sam_account_name",
				Description: "Contains the on-premises SAM account name synchronized from the on-premises directory",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("onPremisesSamAccountName"),
			},
			{
				Name:        "on_premises_security_identifier",
				Description: "Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("onPremisesSecurityIdentifier"),
			},
			{
				Name:        "on_premises_sync_enabled",
				Description: "true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default)",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("onPremisesSyncEnabled"),
			},
			{
				Name:        "photo_entity_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("photo.Entity.additionalData"),
			},
			{
				Name:        "photo_entity_id",
				Description: "Read-only.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("photo.Entity.id"),
			},
			{
				Name:        "photo_height",
				Description: "The height of the photo",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("photo.height"),
			},
			{
				Name:        "photo_width",
				Description: "The width of the photo",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("photo.width"),
			},
			{
				Name:        "planner",
				Description: "Entry-point to Planner resource that might exist for a Unified Group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsPlanner,
			},
			{
				Name:        "preferred_data_location",
				Description: "The preferred data location for the Microsoft 365 group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("preferredDataLocation"),
			},
			{
				Name:        "preferred_language",
				Description: "The preferred language for a Microsoft 365 group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("preferredLanguage"),
			},
			{
				Name:        "proxy_addresses",
				Description: "Email addresses for the group that direct to the same group mailbox",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("proxyAddresses"),
			},
			{
				Name:        "renewed_date_time",
				Description: "Timestamp of when the group was last renewed",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("renewedDateTime"),
			},
			{
				Name:        "security_enabled",
				Description: "Specifies whether the group is a security group",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("securityEnabled"),
			},
			{
				Name:        "security_identifier",
				Description: "Security identifier of the group, used in Windows scenarios",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("securityIdentifier"),
			},
			{
				Name:        "sites",
				Description: "The list of SharePoint sites in this group",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsSites,
			},
			{
				Name:        "team_entity_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.Entity.additionalData"),
			},
			{
				Name:        "team_entity_id",
				Description: "Read-only.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.Entity.id"),
			},
			{
				Name:        "team_channels",
				Description: "The collection of channels & messages associated with the team.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsTeamChannels,
			},
			{
				Name:        "team_classification",
				Description: "An optional label",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.classification"),
			},
			{
				Name:        "team_created_date_time",
				Description: "Timestamp at which the team was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("team.createdDateTime"),
			},
			{
				Name:        "team_description",
				Description: "An optional description for the team",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.description"),
			},
			{
				Name:        "team_display_name",
				Description: "The name of the team.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.displayName"),
			},
			{
				Name:        "teamfun_settings_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.funSettings.additionalData"),
			},
			{
				Name:        "teamfun_settings_allow_custom_memes",
				Description: "If set to true, enables users to include custom memes.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.funSettings.allowCustomMemes"),
			},
			{
				Name:        "teamfun_settings_allow_giphy",
				Description: "If set to true, enables Giphy use.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.funSettings.allowGiphy"),
			},
			{
				Name:        "teamfun_settings_allow_stickers_and_memes",
				Description: "If set to true, enables users to include stickers and memes.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.funSettings.allowStickersAndMemes"),
			},
			{
				Name:        "teamfun_settings_giphy_content_rating",
				Description: "Giphy content rating",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("team.funSettings.giphyContentRating"),
			},
			{
				Name:     "team_group",
				Type:     schema.TypeJSON,
				Resolver: resolveAdGroupsTeamGroup,
			},
			{
				Name:        "teamguest_settings_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.guestSettings.additionalData"),
			},
			{
				Name:        "teamguest_settings_allow_create_update_channels",
				Description: "If set to true, guests can add and update channels.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.guestSettings.allowCreateUpdateChannels"),
			},
			{
				Name:        "teamguest_settings_allow_delete_channels",
				Description: "If set to true, guests can delete channels.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.guestSettings.allowDeleteChannels"),
			},
			{
				Name:        "team_internal_id",
				Description: "A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.internalId"),
			},
			{
				Name:        "team_is_archived",
				Description: "Whether this team is in read-only mode.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.isArchived"),
			},
			{
				Name:        "teammember_settings_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.memberSettings.additionalData"),
			},
			{
				Name:        "teammember_settings_allow_add_remove_apps",
				Description: "If set to true, members can add and remove apps.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowAddRemoveApps"),
			},
			{
				Name:        "teammember_settings_allow_create_private_channels",
				Description: "If set to true, members can add and update private channels.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowCreatePrivateChannels"),
			},
			{
				Name:        "teammember_settings_allow_create_update_channels",
				Description: "If set to true, members can add and update channels.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowCreateUpdateChannels"),
			},
			{
				Name:        "teammember_settings_allow_create_update_remove_connectors",
				Description: "If set to true, members can add, update, and remove connectors.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowCreateUpdateRemoveConnectors"),
			},
			{
				Name:        "teammember_settings_allow_create_update_remove_tabs",
				Description: "If set to true, members can add, update, and remove tabs.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowCreateUpdateRemoveTabs"),
			},
			{
				Name:        "teammember_settings_allow_delete_channels",
				Description: "If set to true, members can delete channels.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.memberSettings.allowDeleteChannels"),
			},
			{
				Name:        "teammessaging_settings_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.messagingSettings.additionalData"),
			},
			{
				Name:        "teammessaging_settings_allow_channel_mentions",
				Description: "If set to true, @channel mentions are allowed.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.messagingSettings.allowChannelMentions"),
			},
			{
				Name:        "teammessaging_settings_allow_owner_delete_messages",
				Description: "If set to true, owners can delete any message.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.messagingSettings.allowOwnerDeleteMessages"),
			},
			{
				Name:        "teammessaging_settings_allow_team_mentions",
				Description: "If set to true, @team mentions are allowed.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.messagingSettings.allowTeamMentions"),
			},
			{
				Name:        "teammessaging_settings_allow_user_delete_messages",
				Description: "If set to true, users can delete their messages.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.messagingSettings.allowUserDeleteMessages"),
			},
			{
				Name:        "teammessaging_settings_allow_user_edit_messages",
				Description: "If set to true, users can edit their messages.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("team.messagingSettings.allowUserEditMessages"),
			},
			{
				Name:        "team_primary_channel",
				Description: "The general channel for the team.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsTeamPrimaryChannel,
			},
			{
				Name:        "team_schedule",
				Description: "The schedule of shifts for this team.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsTeamSchedule,
			},
			{
				Name:        "team_specialization",
				Description: "Optional",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("team.specialization"),
			},
			{
				Name:        "teamtemplate_entity_additional_data",
				Description: "Stores additional data not described in the OpenAPI description found when deserializing",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("team.template.Entity.additionalData"),
			},
			{
				Name:        "teamtemplate_entity_id",
				Description: "Read-only.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.template.Entity.id"),
			},
			{
				Name:        "team_visibility",
				Description: "The visibility of the group and team",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("team.visibility"),
			},
			{
				Name:        "team_web_url",
				Description: "A hyperlink that will go to the team in the Microsoft Teams client",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("team.webUrl"),
			},
			{
				Name:        "theme",
				Description: "Specifies a Microsoft 365 group's color theme",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("theme"),
			},
			{
				Name:        "threads",
				Description: "The group's conversation threads",
				Type:        schema.TypeJSON,
				Resolver:    resolveAdGroupsThreads,
			},
			{
				Name:        "unseen_count",
				Description: "Count of conversations that have received new posts since the signed-in user last visited the group",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("unseenCount"),
			},
			{
				Name:        "visibility",
				Description: "Specifies the group join policy and group content visibility for groups",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("visibility"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "msgraph_ad_group_accepted_senders",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupAcceptedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_app_role_assignments",
				Description: "AppRoleAssignment",
				Resolver:    fetchAdGroupAppRoleAssignments,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "directory_object_entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("DirectoryObject.Entity.additionalData"),
					},
					{
						Name:        "directory_object_entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DirectoryObject.Entity.id"),
					},
					{
						Name:     "directory_object_deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("DirectoryObject.deletedDateTime"),
					},
					{
						Name:        "app_role_id",
						Description: "The identifier (id) for the app role which is assigned to the principal",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("appRoleId"),
					},
					{
						Name:        "created_date_time",
						Description: "The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("createdDateTime"),
					},
					{
						Name:        "principal_display_name",
						Description: "The display name of the user, group, or service principal that was granted the app role assignment",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("principalDisplayName"),
					},
					{
						Name:        "principal_id",
						Description: "The unique identifier (id) for the user, group or service principal being granted the app role",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("principalId"),
					},
					{
						Name:        "principal_type",
						Description: "The type of the assigned principal",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("principalType"),
					},
					{
						Name:        "resource_display_name",
						Description: "The display name of the resource app's service principal to which the assignment is made.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("resourceDisplayName"),
					},
					{
						Name:        "resource_id",
						Description: "The unique identifier (id) for the resource service principal for which the assignment is made",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("resourceId"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_assigned_labels",
				Description: "AssignedLabel",
				Resolver:    fetchAdGroupAssignedLabels,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("additionalData"),
					},
					{
						Name:        "display_name",
						Description: "The display name of the label",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("displayName"),
					},
					{
						Name:        "label_id",
						Description: "The unique identifier of the label.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("labelId"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_assigned_licenses",
				Description: "AssignedLicense",
				Resolver:    fetchAdGroupAssignedLicenses,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("additionalData"),
					},
					{
						Name:        "disabled_plans",
						Description: "A collection of the unique identifiers for plans that have been disabled.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("disabledPlans"),
					},
					{
						Name:        "sku_id",
						Description: "The unique identifier for the SKU.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("skuId"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_lifecycle_policies",
				Description: "GroupLifecyclePolicy",
				Resolver:    fetchAdGroupLifecyclePolicies,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "alternate_notification_emails",
						Description: "List of email address to send notifications for groups without owners",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("alternateNotificationEmails"),
					},
					{
						Name:        "group_lifetime_in_days",
						Description: "Number of days before a group expires and needs to be renewed",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("groupLifetimeInDays"),
					},
					{
						Name:        "managed_group_types",
						Description: "The group type for which the expiration policy applies",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("managedGroupTypes"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_member_of",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_members",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_members_with_license_errors",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupMembersWithLicenseErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_on_premises_provisioning_errors",
				Description: "OnPremisesProvisioningError",
				Resolver:    fetchAdGroupOnPremisesProvisioningErrors,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("additionalData"),
					},
					{
						Name:        "category",
						Description: "Category of the provisioning error",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("category"),
					},
					{
						Name:        "occurred_date_time",
						Description: "The date and time at which the error occurred.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("occurredDateTime"),
					},
					{
						Name:        "property_causing_error",
						Description: "Name of the directory property causing the error",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("propertyCausingError"),
					},
					{
						Name:        "value",
						Description: "Value of the property causing the error.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("value"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_owners",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupOwners,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_permission_grants",
				Description: "ResourceSpecificPermissionGrant",
				Resolver:    fetchAdGroupPermissionGrants,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "directory_object_entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("DirectoryObject.Entity.additionalData"),
					},
					{
						Name:        "directory_object_entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DirectoryObject.Entity.id"),
					},
					{
						Name:     "directory_object_deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("DirectoryObject.deletedDateTime"),
					},
					{
						Name:        "client_app_id",
						Description: "ID of the service principal of the Azure AD app that has been granted access",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("clientAppId"),
					},
					{
						Name:        "client_id",
						Description: "ID of the Azure AD app that has been granted access",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("clientId"),
					},
					{
						Name:        "permission",
						Description: "The name of the resource-specific permission",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("permission"),
					},
					{
						Name:        "permission_type",
						Description: "The type of permission",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("permissionType"),
					},
					{
						Name:        "resource_app_id",
						Description: "ID of the Azure AD app that is hosting the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("resourceAppId"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_photos",
				Description: "ProfilePhoto",
				Resolver:    fetchAdGroupPhotos,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "height",
						Description: "The height of the photo",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("height"),
					},
					{
						Name:        "width",
						Description: "The width of the photo",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("width"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_rejected_senders",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupRejectedSenders,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_settings",
				Description: "GroupSetting",
				Resolver:    fetchAdGroupSettings,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "display_name",
						Description: "Display name of this group of settings, which comes from the associated template.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("displayName"),
					},
					{
						Name:        "template_id",
						Description: "Unique identifier for the template used to create this group of settings",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("templateId"),
					},
					{
						Name:        "values",
						Description: "Collection of name value pairs",
						Type:        schema.TypeJSON,
						Resolver:    resolveAdGroupSettingsValues,
					},
				},
			},
			{
				Name:        "msgraph_ad_group_team_installed_apps",
				Description: "TeamsAppInstallation",
				Resolver:    fetchAdGroupTeamInstalledApps,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "teams_app_entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsApp.Entity.additionalData"),
					},
					{
						Name:        "teams_app_entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsApp.Entity.id"),
					},
					{
						Name:        "teams_app_display_name",
						Description: "The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsApp.displayName"),
					},
					{
						Name:        "teams_app_distribution_method",
						Description: "The method of distribution for the app",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("teamsApp.distributionMethod"),
					},
					{
						Name:        "teams_app_external_id",
						Description: "The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsApp.externalId"),
					},
					{
						Name:        "teams_app_definition_entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.Entity.additionalData"),
					},
					{
						Name:        "teams_app_definition_entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.Entity.id"),
					},
					{
						Name:        "teams_app_definitionbot_entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.bot.Entity.additionalData"),
					},
					{
						Name:        "teams_app_definitionbot_entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.bot.Entity.id"),
					},
					{
						Name:        "teams_app_definitioncreated_by_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.additionalData"),
					},
					{
						Name:        "teams_app_definitioncreated_byapplication_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.application.additionalData"),
					},
					{
						Name:        "teams_app_definitioncreated_byapplication_display_name",
						Description: "The identity's display name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.application.displayName"),
					},
					{
						Name:        "teams_app_definitioncreated_byapplication_id",
						Description: "Unique identifier for the identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.application.id"),
					},
					{
						Name:        "teams_app_definitioncreated_bydevice_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.device.additionalData"),
					},
					{
						Name:        "teams_app_definitioncreated_bydevice_display_name",
						Description: "The identity's display name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.device.displayName"),
					},
					{
						Name:        "teams_app_definitioncreated_bydevice_id",
						Description: "Unique identifier for the identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.device.id"),
					},
					{
						Name:        "teams_app_definitioncreated_byuser_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.user.additionalData"),
					},
					{
						Name:        "teams_app_definitioncreated_byuser_display_name",
						Description: "The identity's display name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.user.displayName"),
					},
					{
						Name:        "teams_app_definitioncreated_byuser_id",
						Description: "Unique identifier for the identity.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.createdBy.user.id"),
					},
					{
						Name:        "teams_app_definition_description",
						Description: "Verbose description of the application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.description"),
					},
					{
						Name:        "teams_app_definition_display_name",
						Description: "The name of the app provided by the app developer.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.displayName"),
					},
					{
						Name:     "teams_app_definition_last_modified_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("teamsAppDefinition.lastModifiedDateTime"),
					},
					{
						Name:        "teams_app_definition_publishing_state",
						Description: "The published status of a specific version of a Teams app",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("teamsAppDefinition.publishingState"),
					},
					{
						Name:        "teams_app_definition_short_description",
						Description: "Short description of the application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.shortDescription"),
					},
					{
						Name:        "teams_app_definition_teams_app_id",
						Description: "The ID from the Teams app manifest.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.teamsAppId"),
					},
					{
						Name:        "teams_app_definition_version",
						Description: "The version number of the application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("teamsAppDefinition.version"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "msgraph_ad_group_team_installed_app_teams_app_app_definitions",
						Description: "TeamsAppDefinition",
						Resolver:    fetchAdGroupTeamInstalledAppTeamsAppAppDefinitions,
						Columns: []schema.Column{
							{
								Name:        "group_team_installed_app_cq_id",
								Description: "Unique CloudQuery ID of msgraph_ad_group_team_installed_apps table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "entity_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("Entity.additionalData"),
							},
							{
								Name:        "entity_id",
								Description: "Read-only.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Entity.id"),
							},
							{
								Name:        "bot_entity_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("bot.Entity.additionalData"),
							},
							{
								Name:        "bot_entity_id",
								Description: "Read-only.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("bot.Entity.id"),
							},
							{
								Name:        "created_by_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("createdBy.additionalData"),
							},
							{
								Name:        "created_byapplication_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("createdBy.application.additionalData"),
							},
							{
								Name:        "created_byapplication_display_name",
								Description: "The identity's display name",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.application.displayName"),
							},
							{
								Name:        "created_byapplication_id",
								Description: "Unique identifier for the identity.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.application.id"),
							},
							{
								Name:        "created_bydevice_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("createdBy.device.additionalData"),
							},
							{
								Name:        "created_bydevice_display_name",
								Description: "The identity's display name",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.device.displayName"),
							},
							{
								Name:        "created_bydevice_id",
								Description: "Unique identifier for the identity.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.device.id"),
							},
							{
								Name:        "created_byuser_additional_data",
								Description: "Stores additional data not described in the OpenAPI description found when deserializing",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("createdBy.user.additionalData"),
							},
							{
								Name:        "created_byuser_display_name",
								Description: "The identity's display name",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.user.displayName"),
							},
							{
								Name:        "created_byuser_id",
								Description: "Unique identifier for the identity.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("createdBy.user.id"),
							},
							{
								Name:        "description",
								Description: "Verbose description of the application.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("description"),
							},
							{
								Name:        "display_name",
								Description: "The name of the app provided by the app developer.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("displayName"),
							},
							{
								Name:     "last_modified_date_time",
								Type:     schema.TypeTimestamp,
								Resolver: schema.PathResolver("lastModifiedDateTime"),
							},
							{
								Name:        "publishing_state",
								Description: "The published status of a specific version of a Teams app",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("publishingState"),
							},
							{
								Name:        "short_description",
								Description: "Short description of the application.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("shortDescription"),
							},
							{
								Name:        "teams_app_id",
								Description: "The ID from the Teams app manifest.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("teamsAppId"),
							},
							{
								Name:        "version",
								Description: "The version number of the application.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("version"),
							},
						},
					},
				},
			},
			{
				Name:        "msgraph_ad_group_team_members",
				Description: "ConversationMember",
				Resolver:    fetchAdGroupTeamMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "display_name",
						Description: "The display name of the user.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("displayName"),
					},
					{
						Name:        "roles",
						Description: "The roles for that user.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("roles"),
					},
					{
						Name:        "visible_history_start_date_time",
						Description: "The timestamp denoting how far back a conversation's history is shared with the conversation member",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("visibleHistoryStartDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_team_operations",
				Description: "TeamsAsyncOperation",
				Resolver:    fetchAdGroupTeamOperations,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:        "attempts_count",
						Description: "Number of times the operation was attempted before being marked successful or failed.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("attemptsCount"),
					},
					{
						Name:        "created_date_time",
						Description: "Time when the operation was created.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("createdDateTime"),
					},
					{
						Name:        "error_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("error.additionalData"),
					},
					{
						Name:        "error_code",
						Description: "Operation error code.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("error.code"),
					},
					{
						Name:        "error_message",
						Description: "Operation error message.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("error.message"),
					},
					{
						Name:        "last_action_date_time",
						Description: "Time when the async operation was last updated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("lastActionDateTime"),
					},
					{
						Name:        "operation_type",
						Description: "Denotes which type of operation is being described.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("operationType"),
					},
					{
						Name:        "status",
						Description: "Operation status.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("status"),
					},
					{
						Name:        "target_resource_id",
						Description: "The ID of the object that's created or modified as result of this async operation, typically a team.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("targetResourceId"),
					},
					{
						Name:        "target_resource_location",
						Description: "The location of the object that's created or modified as result of this async operation",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("targetResourceLocation"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_transitive_member_of",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupTransitiveMemberOfs,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
			{
				Name:        "msgraph_ad_group_transitive_members",
				Description: "DirectoryObject",
				Resolver:    fetchAdGroupTransitiveMembers,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of msgraph_ad_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "entity_additional_data",
						Description: "Stores additional data not described in the OpenAPI description found when deserializing",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Entity.additionalData"),
					},
					{
						Name:        "entity_id",
						Description: "Read-only.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Entity.id"),
					},
					{
						Name:     "deleted_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("deletedDateTime"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAdGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services.Groups
	result, err := svc.Get(nil)
	//result, err := svc.Get(&groups.GroupsRequestBuilderGetOptions{
	//	Q: groups.GroupsRequestBuilderGetQueryParameters{
	//		Expand: //todo add expand object
	//	}
	//})
	if err != nil {
		return err
	}
	for {
		group := result.GetValue()

		cast := make([]serialization.Parsable, len(group))
		for i, v := range group {
			temp := v
			cast[i] = serialization.Parsable(&temp)
			res <- cast[i]
		}

		res <- group
		if result.GetNextLink() == nil {
			break
		}
		result, err = groups.NewGroupsRequestBuilder(*result.GetNextLink(), client.Adapter).Get(nil)
		if err != nil {
			return err
		}
	}
	return nil
}
func resolveAdGroupsCalendar(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	calendar := p.GetCalendar()
	if calendar == nil {
		return nil
	}
	j, err := json.Marshal(calendar)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsCalendarView(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	calendarView := p.GetCalendarView()
	if calendarView == nil {
		return nil
	}
	j, err := json.Marshal(calendarView)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsConversations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	conversations := p.GetConversations()
	if conversations == nil {
		return nil
	}
	j, err := json.Marshal(conversations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsDrive(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	drive := p.GetDrive()
	if drive == nil {
		return nil
	}
	j, err := json.Marshal(drive)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsDrives(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	drives := p.GetDrives()
	if drives == nil {
		return nil
	}
	j, err := json.Marshal(drives)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsEvents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetEvents()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsExtensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetExtensions()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsOnenote(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetOnenote()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsPlanner(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetPlanner()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsSites(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetSites()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsTeamChannels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetTeam()
	if item == nil {
		return nil
	}
	channels := item.GetChannels()
	if channels == nil {
		return nil
	}
	j, err := json.Marshal(channels)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsTeamGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetTeam()
	if item == nil {
		return nil
	}
	group := item.GetGroup()
	if group == nil {
		return nil
	}
	j, err := json.Marshal(group)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsTeamPrimaryChannel(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetTeam()
	if item == nil {
		return nil
	}
	primaryChannel := item.GetPrimaryChannel()
	if primaryChannel == nil {
		return nil
	}
	j, err := json.Marshal(primaryChannel)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsTeamSchedule(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetTeam()
	if item == nil {
		return nil
	}
	schedule := item.GetSchedule()
	if schedule == nil {
		return nil
	}
	j, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func resolveAdGroupsThreads(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", resource.Item)
	}
	item := p.GetThreads()
	if item == nil {
		return nil
	}
	j, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, j)
}
func fetchAdGroupAcceptedSenders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetAcceptedSenders()
	return nil
}
func fetchAdGroupAppRoleAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetAppRoleAssignments()
	return nil
}
func fetchAdGroupAssignedLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetAssignedLabels()
	return nil
}
func fetchAdGroupAssignedLicenses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetAssignedLicenses()
	return nil
}
func fetchAdGroupLifecyclePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetGroupLifecyclePolicies()
	return nil
}
func fetchAdGroupMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetMemberOf()
	return nil
}
func fetchAdGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetMembers()
	return nil
}
func fetchAdGroupMembersWithLicenseErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetMembersWithLicenseErrors()
	return nil
}
func fetchAdGroupOnPremisesProvisioningErrors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetOnPremisesProvisioningErrors()
	return nil
}
func fetchAdGroupOwners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetOwners()
	return nil
}
func fetchAdGroupPermissionGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetPermissionGrants()
	return nil
}
func fetchAdGroupPhotos(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetPhotos()
	return nil
}
func fetchAdGroupRejectedSenders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetRejectedSenders()
	return nil
}
func fetchAdGroupSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetSettings()
	return nil
}
func resolveAdGroupSettingsValues(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(graph.GroupSetting)
	if !ok {
		return fmt.Errorf("expected to have graph.GroupSetting but got %T", resource.Item)
	}

	values := p.GetValues()
	j := map[string]interface{}{}
	for _, v := range values {
		j[*v.GetName()] = *v.GetValue()
	}

	return resource.Set(c.Name, j)
}
func fetchAdGroupTeamInstalledApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	team := p.GetTeam()
	if team == nil {
		return nil
	}
	res <- team.GetInstalledApps()
	return nil
}
func fetchAdGroupTeamInstalledAppTeamsAppAppDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.TeamsAppInstallation)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	app := p.GetTeamsApp()
	if app == nil {
		return nil
	}
	res <- app.GetAppDefinitions()
	return nil
}
func fetchAdGroupTeamMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	team := p.GetTeam()
	if team == nil {
		return nil
	}
	res <- team.GetMembers()
	return nil
}
func fetchAdGroupTeamOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have msgraph.Team but got %T", parent.Item)
	}
	team := p.GetTeam()
	if team == nil {
		return nil
	}

	res <- team.GetOperations()
	return nil
}
func fetchAdGroupTransitiveMemberOfs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetTransitiveMemberOf()
	return nil
}
func fetchAdGroupTransitiveMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(graph.Group)
	if !ok {
		return fmt.Errorf("expected to have graph.Group but got %T", parent.Item)
	}
	res <- p.GetTransitiveMembers()
	return nil
}
