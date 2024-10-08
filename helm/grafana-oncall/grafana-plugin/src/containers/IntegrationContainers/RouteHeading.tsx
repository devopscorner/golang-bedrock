import React from 'react';

import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';
import { Icon, useStyles2 } from '@grafana/ui';

import { LabelBadges } from 'components/LabelsTooltipBadge/LabelsTooltipBadge';
import { RenderConditionally } from 'components/RenderConditionally/RenderConditionally';
import { Text } from 'components/Text/Text';
import { TooltipBadge } from 'components/TooltipBadge/TooltipBadge';
import { ChannelFilter, FilteringTermType } from 'models/channel_filter/channel_filter.types';
import { CommonIntegrationHelper } from 'pages/integration/CommonIntegration.helper';
import { AppFeature } from 'state/features';
import { useStore } from 'state/useStore';

interface RouteHeadingProps {
  className: string;
  routeWording: string;
  routeIndex: number;
  channelFilter: ChannelFilter;
  channelFilterIds: Array<ChannelFilter['id']>;
}

export const RouteHeading: React.FC<RouteHeadingProps> = ({
  className,
  routeWording,
  channelFilterIds,
  channelFilter,
  routeIndex,
}) => {
  const styles = useStyles2(getStyles);

  return (
    <div className={className}>
      <TooltipBadge
        borderType="success"
        text={CommonIntegrationHelper.getRouteConditionWording(channelFilterIds, routeIndex)}
        tooltipTitle={CommonIntegrationHelper.getRouteConditionTooltipWording(
          channelFilterIds,
          routeIndex,
          channelFilter?.filtering_term_type
        )}
        tooltipContent={undefined}
        className={styles.badge}
      />

      {routeWording === 'Default' && <Text type="secondary">Unmatched alerts routed to default route</Text>}
      {routeWording !== 'Default' && <RouteHeadingDisplay channelFilter={channelFilter} />}
    </div>
  );
};

const RouteHeadingDisplay: React.FC<{ channelFilter: ChannelFilter }> = ({ channelFilter }) => {
  const store = useStore();
  const styles = useStyles2(getStyles);
  const hasLabels = store.hasFeature(AppFeature.Labels);

  if (channelFilter?.filtering_term || channelFilter?.filtering_labels) {
    return (
      <>
        <RenderConditionally
          shouldRender={channelFilter.filtering_term_type === FilteringTermType.jinja2 || !hasLabels}
        >
          <Text type="primary" className={styles.routeHeading}>
            {channelFilter.filtering_term}
          </Text>
        </RenderConditionally>

        <RenderConditionally shouldRender={channelFilter.filtering_term_type === FilteringTermType.labels && hasLabels}>
          <LabelBadges labels={channelFilter.filtering_labels} />
        </RenderConditionally>
      </>
    );
  }

  return (
    <>
      <div className={styles.iconExclamation}>
        <Icon name="exclamation-triangle" />
      </div>
      <Text type="primary">Routing template not set</Text>
    </>
  );
};

const getStyles = (theme: GrafanaTheme2) => {
  return {
    badge: css`
      margin-right: 4px;
    `,

    iconExclamation: css`
      color: ${theme.colors.error.main};
    `,

    routeHeading: css`
      max-width: 80%;
      display: block;
      text-overflow: ellipsis;
      overflow: hidden;
    `,
  };
};
